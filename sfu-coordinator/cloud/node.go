package cloud

import "time"

type node struct {
	Ip                string
	Port              string
	PeerCount         int
	Cpu               float64
	isIdle            bool
	lastIdleCheckTime time.Time
}

func (n *node) isCloud(h *Hub) bool {
	cloud := false
	for _, m := range h.machines {
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
		if m.getIP() == n.Ip {
			return &m
		}
	}
	return nil
}
