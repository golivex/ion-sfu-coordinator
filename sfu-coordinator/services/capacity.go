package coordinator

import (
	"sync"
	"time"

	log "github.com/pion/ion-log"
)

type blockedSession struct {
	sync.Mutex
	session map[string]string
}

func (e *etcdCoordinator) startServerAndBlockSession(name string, capacity int) bool {

	e.block.Lock()
	_, ok := e.block.session[name]
	e.block.Unlock()
	if !ok {
		go func() {
			notifyip := make(chan string, 1)
			go e.cloud.StartServerNotify(capacity, name, notifyip)
			log.Infof("waiting for machine ip for session %v", name)
			select {
			case <-time.After(3 * 60 * time.Second):
				log.Infof("time expired for machine ip for session %v", name)
				e.block.Lock()
				delete(e.block.session, name)
				e.block.Unlock()
				return
			case ip := <-notifyip:
				log.Infof("got machine Pip for session %v got ip %v", name, ip)
				e.block.Lock()
				e.block.session[name] = ip
				close(notifyip)
				e.block.Unlock()
				return
				//this will wait till machine has started and ping recieved
				// default:
			}

			log.Infof("got machine ip completed")

		}()
	}
	return ok
}

func (e *etcdCoordinator) isHostBlockedBySession(session string) bool {
	e.block.Lock()
	defer e.block.Unlock()
	_, ok := e.block.session[session]
	return ok
}

func (e *etcdCoordinator) getHostBlockedBySession(session string) *Host {
	e.block.Lock()
	defer e.block.Unlock()
	ip, ok := e.block.session[session]
	if ok {
		if len(ip) == 0 {
			log.Infof("session %v block map %v", session, e.block.session)
			panic("ip cannot be empty here!") //TEMP TODO
		} else {

			for _, host := range e.hosts {
				if host.Ip == ip {
					return &host
				}
			}

			panic("host should be found") //TEMP TODO

			return nil

		}
	} else {
		return nil
	}
}

func (e *etcdCoordinator) ThrottleHost(h *Host) bool {
	if e.cloud != nil {
		cap := e.cloud.GetMachineCapability(h.Ip)
		if cap == -1 {
			log.Infof("unknow capablity for host %v", h.Ip)
		} else {
			log.Infof("capablity for host %v cap %v and current spike is %v", h.Ip, cap, len(h.Spike))
			if len(h.Spike) >= cap/10 {
				log.Infof("10% of capablity allocated recently! need to throttle it")
				return true
			}
		}
	}
	return false
}

func (e *etcdCoordinator) clearBlockedCapacity() {
	for key, host := range e.hosts {
		for session, cap := range host.BlockedCapacity {
			if time.Since(cap.Time) > (60 * time.Second) { //10sec for testing extended this
				log.Infof("clearing capacity blocked for host %v by session %v", host.Ip, session)
				delete(host.BlockedCapacity, session)
			}
		}
		e.hosts[key] = host
	}
}
func (e *etcdCoordinator) getBlockedCapacity(h *Host) int {
	e.clearBlockedCapacity()
	total := 0
	for _, v := range h.BlockedCapacity {
		total = total + v.Cap
	}
	return total
}

func (e *etcdCoordinator) blockHostCapacity(session string, h *Host, cap int) {
	h.BlockedCapacity[session] = capacity{
		Cap:  cap,
		Time: time.Now(),
	}
}
