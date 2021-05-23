package coordinator

import (
	"context"

	log "github.com/pion/ion-log"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const MAX_TRACKS_PER_HOST = 4
const MAX_CLOUD_HOSTS = 0
const MAX_CPU_LOAD = 70

type HostReply struct {
	Host    string
	Status  string
	Session string
}

const NO_HOSTS_RETRY = "NO_HOSTS_RETRY"
const HOST_SESSION_EXISTS = "HOST_SESSION_EXISTS"
const NEW_HOST_ASSIGNED_SESSION = "NEW_HOST_ASSIGNED_SESSION"
const HOST_LOAD_ONLY_SUBSCRIBE = "HOST_LOAD_ONLY_SUBSCRIBE"
const HOST_SCALING = "HOST_SCALING"
const HOST_ALTERNATE = "HOST_ALTERNATE"
const HOST_SCALED_SESSION_FROM_ACTION = "HOST_SCALED_SESSION_FROM_ACTION"
const HOST_FROM_ACTION = "HOST_FROM_ACTION"

type HostSession struct {
	ip   string
	name string
}

func (e *etcdCoordinator) FindHost(session string, isaction bool) HostReply {
	e.mu.Lock()
	defer e.mu.Unlock()

	if len(e.hosts) == 0 {
		return HostReply{
			Status: NO_HOSTS_RETRY,
		}
	}

	sfu.SyncHost(e)

	if sfu.IsScaledSession(session) {
		// hostip := sfu.GetSessionHost(session)
		// here we should find all the hosts which the scaled host has
		// and see which host has avaiblity?
		// but what if there is no host which is avaiable or we have started a new host its not started yet?

		host, opsession := sfu.FindOptmialHost(session, e)
		if host != nil {
			e.SpikeHost(*host)
			return HostReply{
				Host:    host.String(),
				Session: opsession,
				Status:  "SCALED_HOST_OPTIMAL",
			}
		} else {
			nhost := e.findAvailableHost(host)
			if nhost != nil {
				host := sfu.GetOptimalSourceHost(session)
				nsession := sfu.AssignHostToSession(session, nhost, host)

				go MirrorSfu(session, nsession, *host, *nhost)

				return HostReply{
					Status: "HOST_SCALE_MULTIPLE",
				}
			} else {
				return HostReply{
					Status: "SCALED_HOST_NOT_FOUND",
				}
			}
		}
	}

	host := e.checkHostForExistingSession(session)
	status := ""
	if host != nil {
		status = HOST_SESSION_EXISTS
	} else {
		host = e.isHostJustAssignedToSession(session)
		if host != nil {
			status = "SESSION_ASSIGNED_RECENTLY"
		} else {
			host = e.allocateHostToSession(session)
			status = NEW_HOST_ASSIGNED_SESSION
		}
	}
	if e.canHostServe(host) {
		e.SpikeHost(*host)
		return HostReply{
			Host:   host.String(),
			Status: status,
		}
	} else {
		nhost := e.findAvailableHost(host)
		if nhost != nil {
			//need to scale sfu on this host first
			if status == HOST_SESSION_EXISTS {

				nsession := sfu.AssignHostToSession(session, nhost, host)
				go MirrorSfu(session, nsession, *host, *nhost)
				return HostReply{
					// Host:   nhost.String(),
					Status: HOST_SCALING,
				}
			} else {
				e.SpikeHost(*nhost)
				return HostReply{
					Host:   nhost.String(),
					Status: HOST_ALTERNATE,
				}
			}
		} else {
			return HostReply{
				Status: HOST_LOAD_ONLY_SUBSCRIBE,
			}
		}
	}
}

func (e *etcdCoordinator) canHostServe(host *Host) bool {

	log.Infof("checking can host server for %v", *host)

	log.Infof("(host.AudioTracks + host.VideoTracks) %v host %v", (host.AudioTracks + host.VideoTracks), host.String())
	spikepeer, spiketracks, spikecpu := e.GetSpikeLoad(*host)
	log.Infof("Extra spike load %v %v %v", spikepeer, spiketracks, spikecpu)
	if (host.AudioTracks + host.VideoTracks + spiketracks) >= MAX_TRACKS_PER_HOST {
		return false
	} else {
		cpu := host.GetCurrentLoad()
		if (cpu + spikecpu) > MAX_CPU_LOAD {
			return false
		}
		return true
	}
}

func (e *etcdCoordinator) findAvailableHost(existhost *Host) *Host {
	var fhost *Host
	for _, host := range e.hosts {
		if host.String() != existhost.String() {
			if e.canHostServe(&host) {
				log.Infof("host willing to server found %v", host.String())
				fhost = &host
				break
			}
		}
	}
	return fhost
}

func (e *etcdCoordinator) isHostJustAssignedToSession(session string) *Host {
	log.Infof("new session found %v", session)
	kvc := clientv3.NewKV(e.cli)
	gresp, err := kvc.Get(context.Background(), "/temp/"+session)
	var fhost *Host
	if err != nil {
		log.Errorf("host temp key get error", err)
	} else {
		if len(gresp.Kvs) > 0 {
			log.Infof("host just assigned to session")
			hoststr := string(gresp.Kvs[0].Value)
			for _, host := range e.hosts {
				if host.String() == hoststr {
					fhost = &host
					break
				}
			}
		}
	}
	return fhost
}

func (e *etcdCoordinator) allocateHostToSession(session string) *Host {
	var min_load float64
	var min_load_host *Host

	for _, host := range e.hosts {
		cpu := host.GetCurrentLoad()
		log.Infof("host %v loads %v", host.Ip, cpu)
		if cpu > min_load {
			min_load_host = &host
		}
	}

	hoststr := min_load_host.String()
	kvc := clientv3.NewKV(e.cli)
	resp, err := e.cli.Grant(context.Background(), 5)
	if err != nil {
		log.Errorf("lease grant error", err)
	} else {
		_, err = kvc.Put(context.Background(), "/temp/"+session, hoststr, clientv3.WithLease(resp.ID))
		if err != nil {
			log.Errorf("host temp key set error", err)
		}
	}
	return min_load_host
}

func (e *etcdCoordinator) checkHostForExistingSession(session string) *Host {
	var fk *Host
	for key, _ := range e.sessions {
		if e.sessions[key].Name == session {
			sessionhost := e.sessions[key].Host
			sessionport := e.sessions[key].Port
			log.Infof("check host for existing session %v", sessionhost, sessionport)
			for _, host := range e.hosts {
				if host.Ip == sessionhost && host.Port == sessionport {
					fk = &host
					break
				}
			}
			break
		}
	}
	return fk

}
