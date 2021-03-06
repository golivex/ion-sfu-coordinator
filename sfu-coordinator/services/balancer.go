package coordinator

import (
	"context"

	log "github.com/pion/ion-log"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const CAN_HOST_MIRROR = false

//as of now host cannot scale, first i am implemeting simply strategy
//i.e if host cannot scale then upto a specific load hosts can join
// after that load hosts can join only as a subscriber
// after that load hosts cannot join at all
// hopefully till then feat relay is completed in ionsfu

const MAX_TRACKS_PER_HOST = 9999      //if more than X tracks cannot connect to host at all
const MAX_PUBLISH_TRACK_THRESH = 9999 //if more than X tracks cannot publish anymore, can only subscribe

const MAX_CPULOAD_PER_HOST = 80       // if more than X cpu load, cannot connect to the host at all until load goes down, also load goes down over a period of time
const MAX_PUBLISH_CPULOAD_THRESH = 70 // if more than X cpu, cannot publish can only subscribe

type sessionStat struct {
	PeerCount   int
	AudioTracks int
	VideoTracks int
}
type HostReply struct {
	Host    string
	Domain  string
	Status  string
	Session string
	Publish bool
	Stats   sessionStat
}

const NO_HOSTS_RETRY = "NO_HOSTS_RETRY"
const HOST_SESSION_EXISTS = "HOST_SESSION_EXISTS"
const NEW_HOST_ASSIGNED_SESSION = "NEW_HOST_ASSIGNED_SESSION"
const HOST_LOAD_ONLY_SUBSCRIBE = "HOST_LOAD_ONLY_SUBSCRIBE"
const HOST_SCALING = "HOST_SCALING"
const HOST_ALTERNATE = "HOST_ALTERNATE"
const HOST_SCALED_SESSION_FROM_ACTION = "HOST_SCALED_SESSION_FROM_ACTION"
const HOST_FROM_ACTION = "HOST_FROM_ACTION"

const DEFAULT_CAPACITY = 5

type HostSession struct {
	ip   string
	name string
}

func (e *etcdCoordinator) FindHost(session string, capacity int, role string) HostReply {
	e.mu.Lock()
	defer e.mu.Unlock()

	if role == "" {
		role = "pubsub"
		//allowed roles is pubsub/sub/pub
	}
	if capacity == -1 {
		capacity = DEFAULT_CAPACITY
	} else {
		log.Infof("got capacity %v", capacity)
	}
	if role == "sub" {
		capacity = capacity / 10
	}

	if len(e.hosts) == 0 {
		if e.cloud != nil {
			if e.isHostBlockedBySession(session) {
				return HostReply{
					Status: "SERVER_PROVISIONING",
				}
			} else {
				if e.cloud.CanAddMachine() {
					if !e.startServerAndBlockSession(session, capacity) {
						return HostReply{
							Status: "STARTING_NEW_CLOUD_SERVER",
						}
					} else {
						return HostReply{
							Status: "SERVER_PROVISIONING",
						}
					}
				} else {
					return HostReply{
						Status: "NO_HOSTS_CANANOT_START_CLOUD",
					}
				}
			}
		} else {
			return HostReply{
				Status: NO_HOSTS_RETRY,
			}
		}

	}

	// e.debugSession()
	// log.Infof("before find host session debug")

	// if CAN_HOST_MIRROR {
	// 	sfu.SyncHost(e)

	// 	if sfu.IsScaledSession(session) {
	// 		// hostip := sfu.GetSessionHost(session)
	// 		// here we should find all the hosts which the scaled host has
	// 		// and see which host has avaiblity?
	// 		// but what if there is no host which is avaiable or we have started a new host its not started yet?

	// 		host, opsession := sfu.FindOptmialHost(session, e)
	// 		if host != nil {
	// 			e.SpikeHost(host)
	// 			if e.ThrottleHost(host) {
	// 				return HostReply{
	// 					Status: "HOST_THROTTLE",
	// 				}
	// 			} else {
	// 				return HostReply{
	// 					Host:    host.String(),
	// 					Session: opsession,
	// 					Status:  "SCALED_HOST_OPTIMAL",
	// 				}
	// 			}

	// 		} else {
	// 			nhost := e.findAvailableHost(nil) //TODO extend this to pass arary of all hosts which are already scaled
	// 			if nhost != nil {
	// 				host := sfu.GetOptimalSourceHost(session)
	// 				nsession := sfu.AssignHostToSession(session, nhost, host)

	// 				go MirrorSfu(session, nsession, *host, *nhost)

	// 				return HostReply{
	// 					Status: "HOST_SCALE_MULTIPLE",
	// 				}
	// 			} else {
	// 				return HostReply{
	// 					Status: "SCALED_HOST_NOT_FOUND",
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	host := e.checkHostForExistingSession(session)
	status := ""
	if host != nil {
		status = HOST_SESSION_EXISTS
	} else {
		host = e.isHostJustAssignedToSession(session)
		if host != nil {
			status = "SESSION_ASSIGNED_RECENTLY"
		} else {
			if e.isHostBlockedBySession(session) {
				return HostReply{
					Status: "SERVER_PROVISIONING",
				}
			} else {
				host = e.allocateHostToSession(session, capacity)
				if host == nil {
					if e.cloud != nil {
						if e.cloud.CanAddMachine() {
							if !e.startServerAndBlockSession(session, capacity) {
								return HostReply{
									Status: "HOST_LOAD_EXECEEDED_SCALING_NEW",
								}
							} else {
								return HostReply{
									Status: "HOST_LOAD_EXECEEDED_SCALING_NEW_PROVISIONING",
								}
							}
						} else {
							return HostReply{
								Status: "HOST_UNAVAILABLE",
							}
						}
					} else {
						return HostReply{
							Status: "HOST_UNAVAILABLE",
						}
					}

				} else {
					status = NEW_HOST_ASSIGNED_SESSION
				}
			}
		}
	}
	if canServe, canPublish := e.canHostServe(host, role); canServe {
		log.Infof("canServe %v canPublish %v", canServe, canPublish)
		if role == "pubsub" || role == "pub" {
			e.blockHostCapacity(session, host, capacity)
		}
		if e.ThrottleHost(host, role) {
			return HostReply{
				Status: "HOST_THROTTLE",
			}
		} else {
			e.SpikeHost(host, role)
			stats := e.getSessionStats(host, session)
			return HostReply{
				Host:    host.String(),
				Domain:  host.Domain,
				Status:  status,
				Publish: canPublish,
				Stats:   stats,
			}
		}
	} else {
		// if CAN_HOST_MIRROR {
		// 	nhost := e.findAvailableHost(host)
		// 	if nhost != nil && CAN_HOST_MIRROR {
		// 		//need to scale sfu on this host first
		// 		if status == HOST_SESSION_EXISTS {

		// 			nsession := sfu.AssignHostToSession(session, nhost, host)
		// 			go MirrorSfu(session, nsession, *host, *nhost)
		// 			return HostReply{
		// 				Host:    nhost.String(),
		// 				Status:  HOST_SCALING,
		// 				Session: nsession,
		// 			}
		// 		} else {
		// 			e.SpikeHost(nhost)
		// 			e.blockHostCapacity(session, host, capacity)
		// 			return HostReply{
		// 				Host:   nhost.String(),
		// 				Status: HOST_ALTERNATE,
		// 			}
		// 		}
		// 	}
		// }
		log.Infof("host load execeeded host cannot serve")
		return HostReply{
			Status: "HOST_LOAD_EXCEED",
		}

	}
}

func (e *etcdCoordinator) getSessionStats(host *Host, session string) sessionStat {
	for key, _ := range e.sessions {
		if e.sessions[key].Name == session {
			sessionhost := e.sessions[key].Host
			sessionport := e.sessions[key].Port
			if host.Ip == sessionhost && host.Port == sessionport {
				return sessionStat{
					PeerCount:   e.sessions[key].PeerCount,
					AudioTracks: e.sessions[key].AudioTracks,
					VideoTracks: e.sessions[key].VideoTracks,
				}
			}
		}
	}
	return sessionStat{}
}

func (e *etcdCoordinator) canHostServe(host *Host, role string) (canServe bool, canPublish bool) {

	log.Debugf("checking can host server for %v", *host)

	log.Infof("(host.AudioTracks + host.VideoTracks) %v host %v host.GetCurrentLoad() %v", (host.AudioTracks + host.VideoTracks), host.String(), host.GetCurrentLoad())
	spikepeer, spiketracks, spikecpu := e.GetSpikeLoad(*host)
	cpu := host.GetCurrentLoad()
	log.Infof("Extra spike load %v %v %v", spikepeer, spiketracks, spikecpu)

	if (cpu + spikecpu) > MAX_CPULOAD_PER_HOST {
		log.Infof("cpu load already at max host cannot serve real load %v cpu %v  max load %v", cpu, (cpu + spikecpu), MAX_CPULOAD_PER_HOST)
		return false, false
	} else {

		if (cpu + spikecpu) > MAX_PUBLISH_CPULOAD_THRESH {
			log.Infof("cpu load above publish threashhold cpu %v load %v", (cpu + spikecpu), MAX_PUBLISH_CPULOAD_THRESH)
			return true, false
		} else {

			trackload := (host.AudioTracks + host.VideoTracks + spiketracks)
			if trackload >= MAX_TRACKS_PER_HOST && (role == "pubsub" || role == "pub") {
				log.Infof("max tracks reached for this host tracks %v max tracks per host", trackload, MAX_TRACKS_PER_HOST)
				return false, false
			} else {

				if trackload >= MAX_PUBLISH_TRACK_THRESH && (role == "pubsub" || role == "pub") {
					log.Infof("max tracks reached for publishing host %v max tracks per host", trackload, MAX_TRACKS_PER_HOST)
					return true, false
				} else {
					log.Infof("host can serve and publish")
					return true, true
				}

			}

		}

	}
}

func (e *etcdCoordinator) findAvailableHost(existhost *Host, role string) *Host {
	var fhost *Host
	if existhost == nil {
		for _, host := range e.hosts {
			if canServe, canPublish := e.canHostServe(&host, role); canServe {
				log.Infof("canServe %v canPublish %v", canServe, canPublish)
				log.Infof("host willing to server found %v", host.String())
				fhost = &host
				break
			}
		}
	} else {
		for _, host := range e.hosts {
			if host.String() != existhost.String() {
				if canServe, canPublish := e.canHostServe(&host, role); canServe {
					log.Infof("canServe %v canPublish %v", canServe, canPublish)
					log.Infof("host willing to server found %v", host.String())
					fhost = &host
					break
				}
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

func (e *etcdCoordinator) allocateHostToSession(session string, capacity int) *Host {
	min_load := float64(99999)
	max_load := float64(MAX_PUBLISH_CPULOAD_THRESH)
	var min_load_host *Host

	blockedhost := e.getHostBlockedBySession(session)
	if blockedhost == nil {
		for _, host := range e.hosts {

			//first check capablity
			if e.cloud != nil && capacity != -1 {
				hcap := e.cloud.GetMachineCapability(host.Ip)
				blocked := e.getBlockedCapacity(&host)
				if hcap == -1 {
					log.Infof("unknow host capablity, assiming that host will be able to handle the load %v", hcap)
				} else if capacity > (hcap - blocked) {
					log.Infof("skipping host as machine cannot handle the required capablity %v avaiablble capacity %v blocked capacvity %v", capacity, hcap, blocked)
					continue
				}

			}

			cpu := host.GetCurrentLoad()
			log.Infof("host %v loads %v", host.Ip, cpu)
			if cpu <= min_load {
				min_load_host = &host
				min_load = cpu
			}
		}
		if min_load > max_load {
			min_load_host = nil

		}
	} else {
		log.Infof("blocked host found for session %v host %v", session, blockedhost.String())
		min_load_host = blockedhost
	}
	if min_load_host != nil {
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
