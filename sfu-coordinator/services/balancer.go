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
	Host   string
	Status string
}

const NO_HOSTS_RETRY = "NO_HOSTS_RETRY"
const HOST_SESSION_EXISTS = "HOST_SESSION_EXISTS"
const HOST_ASSIGNED_SESSION = "HOST_ASSIGNED_SESSION"
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

	log.Infof("is action %v", isaction)

	if len(e.hosts) == 0 {
		return HostReply{
			Status: NO_HOSTS_RETRY,
		}
	}

	if isaction {
		if sfu.IsScaledSession(session) {
			hostip := sfu.GetSessionHost(session)
			if !hostip.Empty() {
				log.Infof("empty scaled host")
				return HostReply{
					Host:   hostip.String(),
					Status: HOST_SCALED_SESSION_FROM_ACTION,
				}
			}
		}
	}

	host, ok := e.checkHostForExistingSession(session)
	status := HOST_SESSION_EXISTS
	if !ok {
		host = e.allocateHostToSession(session)
		status = HOST_ASSIGNED_SESSION
	}
	if isaction {
		//not checking for load etc from called from actions
		return HostReply{
			Host:   host.String(),
			Status: HOST_FROM_ACTION,
		}
	}
	if e.canHostServe(host) {
		return HostReply{
			Host:   host.String(),
			Status: status,
		}
	} else {
		nhost, ok := e.findAvailableHost(host)
		if ok {
			//need to scale sfu on this host first
			if status == HOST_SESSION_EXISTS {

				nsession := sfu.AssignHostToSession(session, nhost)

				go MirrorSfu(session, nsession, host.Ip)

				return HostReply{
					Host:   nhost.String(),
					Status: HOST_SCALING,
				}
			} else {

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

func (e *etcdCoordinator) canHostServe(host Host) bool {
	log.Infof("(host.AudioTracks + host.VideoTracks) %v host %v", (host.AudioTracks + host.VideoTracks), host)
	if (host.AudioTracks + host.VideoTracks) >= MAX_TRACKS_PER_HOST {
		return false
	} else {
		cpu := host.GetCurrentLoad()
		if cpu > MAX_CPU_LOAD {
			return false
		}
		return true
	}
}

func (e *etcdCoordinator) findAvailableHost(existhost Host) (Host, bool) {
	var fhost Host
	for _, host := range e.hosts {
		if host.String() != existhost.String() {
			if e.canHostServe(host) {
				fhost = host
				break
			}
		}
	}
	if fhost.Empty() {
		return fhost, false
	} else {
		return fhost, true
	}

}

func (e *etcdCoordinator) allocateHostToSession(session string) Host {
	log.Infof("new session found %v", session)
	kvc := clientv3.NewKV(e.cli)
	gresp, err := kvc.Get(context.Background(), "/temp/"+session)
	if err != nil {
		log.Errorf("host temp key get error", err)
	} else {
		if len(gresp.Kvs) > 0 {
			log.Infof("host just assigned to session")
			hoststr := string(gresp.Kvs[0].Value)
			var fhost Host
			for _, host := range e.hosts {
				if host.String() == hoststr {
					fhost = host
					break
				}
			}
			if !fhost.Empty() {
				return fhost
			}
		}
	}

	var min_load float64
	var min_load_host Host

	for _, host := range e.hosts {
		cpu := host.GetCurrentLoad()
		log.Infof("host %v loads %v", host.Ip, host.Loads)
		if cpu > min_load {
			min_load_host = host
		}
	}

	hoststr := min_load_host.String()

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

func (e *etcdCoordinator) checkHostForExistingSession(session string) (Host, bool) {
	var fk Host
	for key, _ := range e.sessions {
		if e.sessions[key].Name == session {
			sessionhost := e.sessions[key].Host
			sessionport := e.sessions[key].Port
			log.Infof("check host for existing session %v", sessionhost, sessionport)
			for _, host := range e.hosts {
				if host.Ip == sessionhost && host.Port == sessionport {
					fk = host
				}
			}
			break
		}
	}
	if !fk.Empty() {
		return fk, true
	} else {

		return fk, false
	}

}
