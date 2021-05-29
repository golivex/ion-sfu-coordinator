package coordinator

import (
	log "github.com/pion/ion-log"
)

func (e *etcdCoordinator) getReadyActionHost() *Host {
	var h *Host
	for _, host := range e.actionhosts {
		if host.Tasks == 0 {
			if e.checkActionNode(host.Ip, host.Port) {
				h = &host
			}
		}
	}
	return h
}

func (e *etcdCoordinator) getActionHostByIp(ip string) *Host {
	var h *Host
	for _, host := range e.actionhosts {
		if host.Ip == ip {
			h = &host
		}
	}
	return h
}

func (e *etcdCoordinator) startActionHost(capacity int) chan string {
	notifyip := make(chan string, 1)
	log.Infof("startActionHost %v", capacity)
	go func() {
		log.Infof("starting action machine with capacity %v", capacity)
		if e.cloud.StartActionServerNotify(capacity, notifyip) {
			log.Infof("waiting for ip of action machine")
		} else {
			log.Infof("unable to start action machine")
		}
	}()
	return notifyip
}
