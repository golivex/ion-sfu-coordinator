package cloud

import (
	"time"

	log "github.com/pion/ion-log"
)

type node struct {
	Ip                string
	Port              string
	PeerCount         int
	Cpu               float64
	isIdle            bool
	lastIdleCheckTime time.Time
	lastPing          time.Time
}

func (n *node) isCloud(h *Hub) bool {
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

func (n *node) checkAllNodeIdle(h *Hub) bool {
	all_idle := true
	for _, n2 := range h.nodes {
		if n2.Ip == n.Ip && !n2.isIdle {
			all_idle = false
		}
	}
	return all_idle
}

func (n *node) getCloudMachine(h *Hub) *machine {
	for _, m := range h.machines {
		log.Infof("checking for cloud machine from node m.getIp %v n.Ip%v", m.getIP(), n.Ip)
		if m.getIP() == n.Ip {
			return &m
		}
	}
	return nil
}

func (h *Hub) checkIdleNodes() {
	h.Lock()
	defer h.Unlock()
	// log.Infof("checking idle nodes %v", len(h.nodes))
	for idx, n := range h.nodes {
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
					// log.Infof("all nodes on this ip are not idle so cannot delete this server %v", n.Ip)
				}
			}

		}

		if n.PeerCount == 0 {
			// log.Infof("node is idle %v %v", n.Ip, n.Port)
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
