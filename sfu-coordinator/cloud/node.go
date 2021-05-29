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
