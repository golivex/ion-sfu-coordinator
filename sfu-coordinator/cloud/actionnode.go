package cloud

import (
	"time"

	log "github.com/pion/ion-log"
)

type actionnode struct {
	Ip                string
	Port              string
	Tasks             int
	Cpu               float64
	isIdle            bool
	lastIdleCheckTime time.Time
	lastPing          time.Time
}

func (n *actionnode) isCloud(h *Hub) bool {
	cloud := false
	log.Infof("checking machines %v", len(h.machines))
	for _, m := range h.machines {
		log.Infof("m get ip %v nip %v", m.getIP(), n.Ip)
		if m.getIP() == n.Ip {
			cloud = true
		}
	}
	return cloud
}

func (n *actionnode) checkAllNodeIdle(h *Hub) bool {
	all_idle := true
	for _, n2 := range h.nodes {
		if n2.Ip == n.Ip && !n2.isIdle {
			all_idle = false
		}
	}
	return all_idle
}

func (n *actionnode) getCloudMachine(h *Hub) *machine {
	for _, m := range h.machines {
		log.Infof("checking for cloud machine from node m.getIp %v node.Ip%v", m.getIP(), n.Ip)
		if m.getIP() == n.Ip {
			return &m
		}
	}
	return nil
}

func (h *Hub) checkIdleActionNodes() {
	h.Lock()
	defer h.Unlock()
	log.Infof("checking idle action nodes %v", len(h.actionnodes))
	for idx, n := range h.actionnodes {
		if n.isIdle {

			if time.Since(n.lastIdleCheckTime) > (IDLE_TIMEOUT_CLOUD_HOST * time.Second) {

				//check if all nodes are idle for this specific ip
				// as we can have multiple nodes on a single ip
				all_idle := n.checkAllNodeIdle(h)
				if all_idle {
					//check if its a cloud instance
					if n.isCloud(h) {
						log.Infof("node is idle after %v sec delete it as its cloud instance", IDLE_TIMEOUT_CLOUD_HOST)
						m := n.getCloudMachine(h)
						if m != nil {
							if len(h.machines) > MINIMUM_CLOUD_HOSTS {
								log.Infof("deleting host %v", m.getIP())
								go DeleteInstance(*m)
								//TODO need to see how we can delete avaiable host here instance maybe even use etcd?
								delete(h.machines, m.Id)
							} else {
								log.Infof("cannot delete cloud instance as minimum of %v instances required", MINIMUM_CLOUD_HOSTS)
							}
						}
					} else {
						log.Infof("node is idle after %v sec but its not a cloud instance", IDLE_TIMEOUT_CLOUD_HOST)
					}

				} else {
					log.Infof("all nodes on this ip are not idle so cannot delete this server %v", n.Ip)
				}
			}

		}

		if n.Tasks == 0 {
			log.Infof("action node is idle %v %v tasks %v", n.Ip, n.Port, n.Tasks)
			if !n.isIdle {
				n.lastIdleCheckTime = time.Now()
			}
			n.isIdle = true
		} else {
			n.isIdle = false
		}
		h.actionnodes[idx] = n
	}
}

func (h *Hub) checkDeadActionNodes() {
	h.Lock()
	defer h.Unlock()
	//check dead nodes which are not clouds instances
	for idx, n := range h.actionnodes {
		if n.getCloudMachine(h) == nil {
			log.Infof("node %v is not a cloud instance", n.Ip)

			if time.Since(n.lastPing) > 15*time.Second {
				log.Infof("removing dead node %v port %v", n.Ip, n.Port)
				h.actionnodes = append(h.actionnodes[:idx], h.actionnodes[idx+1:]...)
				break
			}
		}
	}
}

func (h *Hub) StartActionServerNotify(capacity int, notify chan<- string) bool {

	m, err := StartInstance(capacity, -1, true)
	if err != nil {
		log.Errorf("unable to start server %v", err)
		return false
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
		return true
	}

}
