package cloud

import (
	"context"
	"sync"
	"time"

	log "github.com/pion/ion-log"
)

var CAPABLITY = map[string]int{
	"5.9.18.28": 200,
}

type ping struct {
	isDead          bool
	lastIsDeadCheck time.Time
}

type Hub struct {
	sync.Mutex
	machines map[string]machine //machine are mainly cloud based instances
	nodes    []node             // nodes are sfu nodes running on cloud and normal server instances, one server can have multiple nodes as well

	machinePingMap map[string]ping

	cloudOp            bool // is cloud operation in progress
	lastMachineStarted map[string]machineOnline
}

type machineOnline struct {
	time         time.Time
	shouldnotify bool
	notify       chan<- string
}

func RegisterHub(ctx context.Context) *Hub {
	h := &Hub{
		machines:           make(map[string]machine),
		nodes:              []node{},
		machinePingMap:     make(map[string]ping),
		lastMachineStarted: make(map[string]machineOnline),
	}
	go h.autoScaleNodes(ctx)
	go h.syncCloudMachines()

	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				h.syncCloudMachines()
			}
		}
	}()

	return h
}

const IDLE_TIMEOUT_CLOUD_HOST = 60
const WAIT_TIMEOUT_DELETE_CLOUD_DEAD = 15
const MAX_MACHINE_LOAD = 70
const MINIMUM_CLOUD_HOSTS = 0
const MAX_CLOUD_HOSTS = 2

func (h *Hub) startDefaultServer() {
	h.Lock()
	cloudOp := h.cloudOp
	h.Unlock()
	// need to wait here for the server to come online before we start another server
	if len(h.lastMachineStarted) != 0 {
		log.Infof("waiting for last machine started to get online")
	} else {
		if cloudOp {
			log.Infof("cloud op already in progress so skipping!")
		} else {
			h.Lock()
			h.cloudOp = true
			h.Unlock()
			go func() {
				m, err := StartInstance(-1, -1)
				if err != nil {
					log.Errorf("unable to start server %v", err)
				} else {
					h.Lock()
					h.machines[m.Id] = m
					h.lastMachineStarted[m.getIP()] = machineOnline{
						time:         time.Now(),
						shouldnotify: false,
					}
					time.AfterFunc(2*60*time.Second, func() {
						h.Lock()
						log.Infof("machine timeout starteding..... deleting it from here", m.getIP())
						delete(h.lastMachineStarted, m.getIP())
						h.Unlock()
					})
					h.Unlock()
				}
				h.Lock()
				h.cloudOp = false
				h.Unlock()
			}()
		}
	}
}
func (h *Hub) StartServerNotify(capacity int, session string, notify chan<- string) {
	h.Lock()
	h.cloudOp = true
	h.Unlock()
	m, err := StartInstance(capacity, -1)
	if err != nil {
		log.Errorf("unable to start server %v", err)
	} else {
		h.Lock()
		h.machines[m.Id] = m
		h.lastMachineStarted[m.getIP()] = machineOnline{
			time:         time.Now(),
			shouldnotify: true,
			notify:       notify,
		}
		h.Unlock()
		// notify <- m.getIP() this is wrong. we are doing notify when we get ping from machine
		time.AfterFunc(2*60*time.Second, func() {
			h.Lock()
			log.Infof("machine timeout starteding..... deleting it from here", m.getIP())
			delete(h.lastMachineStarted, m.getIP())
			h.Unlock()
		})
	}
	h.Lock()
	h.cloudOp = false
	h.Unlock()

}

func (h *Hub) syncCloudMachines() {
	h.Lock()
	defer h.Unlock()
	log.Infof("sync cloud machines, existing machines %v", len(h.machines))
	machines := GetInstanceList()
	for _, m := range machines {
		_, ok := h.machines[m.Id]
		if !ok {
			log.Infof("new machine added %v created time %v", m.getIP(), m.CreationTimestamp)
			h.machines[m.Id] = m
		}
	}

	for id, m := range h.machines {
		found := false
		for _, exm := range machines {
			log.Infof("machine found from get instance id %v ip %v machine type %v", id, exm.getIP(), exm.getInstanceType())
			if exm.Id == id {
				found = true
				break
			}
		}
		if !found {
			log.Infof("machine deleted %v", m.getIP())
			delete(h.machines, id)
		} else {
			log.Infof("machine already exists %v", m.getIP())
		}

	}
}

func (h *Hub) checkIdleNodes() {
	h.Lock()
	defer h.Unlock()
	log.Infof("checking idle nodes %v", len(h.nodes))
	for idx, n := range h.nodes {
		if n.isIdle {

			if time.Since(n.lastIdleCheckTime) > (IDLE_TIMEOUT_CLOUD_HOST * time.Second) {

				//check if all nodes are idle for this specific ip
				// as we can have multiple nodes on a single ip
				all_idle := n.checkAllNodeIdle(h)
				if all_idle {
					//check if its a cloud instance
					if n.isCloud(h) {
						log.Infof("node is idle after 20 sec delete it as its cloud instance")
						m := n.getCloudMachine(h)
						if m != nil {
							if len(h.machines) > MINIMUM_CLOUD_HOSTS {
								go DeleteInstance(*m)
								delete(h.machines, m.Id)
							} else {
								log.Infof("cannot delete cloud instance as minimum of %v instances required", MINIMUM_CLOUD_HOSTS)
							}
						}
					} else {
						log.Infof("node is idle after 20 sec but its not a cloud instance")
					}

				} else {
					log.Infof("all nodes on this ip are not idle so cannot delete this server %v", n.Ip)
				}
			}

		}

		if n.PeerCount == 0 {
			log.Infof("node is idle %v %v", n.Ip, n.Port)
			if !n.isIdle {
				n.lastIdleCheckTime = time.Now()
			}
			n.isIdle = true
		} else {
			n.isIdle = false
		}
		h.nodes[idx] = n
	}
}

func (h *Hub) checkDeadNodes() {
	h.Lock()
	defer h.Unlock()
	//check dead nodes which are not clouds instances
	for idx, n := range h.nodes {
		if n.getCloudMachine(h) == nil {
			log.Infof("node %v is not a cloud instance", n.Ip)

			if time.Since(n.lastPing) > 15*time.Second {
				log.Infof("removing dead node %v port %v", n.Ip, n.Port)
				h.nodes = append(h.nodes[:idx], h.nodes[idx+1:]...)
				break
			}
		}
	}
}

func (h *Hub) checkDeadMachines() {
	h.Lock()
	defer h.Unlock()
	log.Infof("checking dead machines %v", len(h.machines))
	for id, m := range h.machines {

		found := false
		for _, n := range h.nodes {
			if n.Ip == m.getIP() {
				found = true
				log.Infof("ip found machine is not dead %v", m.getIP())
				break
			}
		}

		if !found {
			log.Infof("dead machine %v", m.getIP())
			p, ok := h.machinePingMap[id]
			if !ok {
				log.Infof("adding ping %v", id)
				h.machinePingMap[id] = ping{
					isDead:          true,
					lastIsDeadCheck: time.Now(),
				}
			} else {
				log.Infof("updating from last poing %v", p)
				if time.Since(m.CreationTimestamp) > (2 * 60 * time.Second) {
					if time.Since(p.lastIsDeadCheck) > (WAIT_TIMEOUT_DELETE_CLOUD_DEAD * time.Second) {
						log.Infof("delete this cloud machine as its dead %v", m.getIP())
						go DeleteInstance(m)
						delete(h.machinePingMap, id)
					} else {
						log.Infof("delete this machine %v but waiting for %v seconds", m.getIP(), WAIT_TIMEOUT_DELETE_CLOUD_DEAD)
					}
				} else {
					log.Infof("machine creating under 2min", time.Since(m.CreationTimestamp))
				}

			}
		} else {
			_, ok := h.machinePingMap[id]
			if ok {
				log.Infof("machine is not dead anymore %v", id)
				delete(h.machinePingMap, id)
			}
		}

	}
}

func (h *Hub) scaleNodeLoad() {
	h.Lock()
	defer h.Unlock()

	log.Infof("checking loads across nodes")

	if len(h.machines) < MINIMUM_CLOUD_HOSTS {
		log.Infof("minimum machines need not met %v starting hosts", MINIMUM_CLOUD_HOSTS)
		go h.startDefaultServer()
	}

	has_node_unload := false
	for _, n := range h.nodes {

		log.Infof("node %v port %v load %v peers %v", n.Ip, n.Port, n.Cpu, n.PeerCount)
		if n.Cpu < float64(MAX_MACHINE_LOAD) {
			has_node_unload = true
		}

	}

	if !has_node_unload && len(h.nodes) != 0 {
		log.Infof("start new server as all machines are above 70 per load")
		go h.startDefaultServer()
	}
}

func (h *Hub) autoScaleNodes(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ctx.Done():
			log.Infof("closing auto scale nodes")
			ticker.Stop()
			return

		case <-ticker.C:
			log.Infof("auto scaling nodes")
			h.checkDeadNodes()
			h.checkDeadMachines()
			h.checkIdleNodes()
			h.scaleNodeLoad()
		}
	}

}

func (h *Hub) DeleteNode(ip string, port string) {
	for idx, n := range h.nodes {
		if n.Ip == ip && n.Port == port {
			h.nodes = append(h.nodes[:idx], h.nodes[idx+1:]...)
			break
		}
	}
}

func (hub *Hub) UpdateNodeLoad(ip string, port string, peer int, cpu float64) {

	log.Infof("updating host load ip%v port%v peer %v cpu%v", ip, port, peer, cpu)

	if len(hub.lastMachineStarted) > 0 {
		online, ok := hub.lastMachineStarted[ip]
		if ok {
			log.Infof("last machine started is online! %v took time %v", ip, time.Since(online.time))
			if online.shouldnotify {
				online.notify <- ip
			}
			delete(hub.lastMachineStarted, ip)
		}
	}
	found := false
	for idx, n := range hub.nodes {

		if n.Ip == ip && n.Port == port {
			n.PeerCount = peer
			n.Cpu = cpu
			n.lastPing = time.Now()
			hub.nodes[idx] = n
			found = true
			break
		}
	}
	if !found {
		hub.nodes = append(hub.nodes, node{
			Ip:        ip,
			Port:      port,
			PeerCount: peer,
			Cpu:       cpu,
			lastPing:  time.Now(),
		})
	}
}

func (h *Hub) CanAddMachine() bool {
	em := len(h.machines) + len(h.lastMachineStarted)
	if h.cloudOp {
		em = em + 1
	}

	return em < MAX_CLOUD_HOSTS
}

func (h *Hub) GetMachineCapability(ip string) int {
	cap, ok := CAPABLITY[ip]
	if ok {
		return cap
	} else {

		for _, m := range h.machines {
			log.Infof("checking capablity with machines getIp %v ip %v instance type %v", m.getIP(), ip, m.getInstanceType())
			if m.getIP() == ip {
				return GetInstanceCapablity(m.getInstanceType())
			}
		}
	}
	return -1
}

// func Test() {
// 	_, err := StartInstance(-1)
// 	if err != nil {
// 		log.Infof("err %v", err)
// 	} else {
// 		log.Infof("new machine started")
// 	}
// 	machines := GetInstanceList()
// 	log.Infof("existing machines %v", len(machines))
// 	err = DeleteInstance(machines[0])
// 	if err != nil {
// 		log.Infof("error deleting instance %v", err)
// 	}
// 	os.Exit(1)
// }
