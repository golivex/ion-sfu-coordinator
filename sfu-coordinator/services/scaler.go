package coordinator

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"

	log "github.com/pion/ion-log"
)

type SfuScaler struct {
	hostmap map[string]*Host //this will map all scaled session and their assigned hosts
	mu      sync.Mutex
}

var sfu = SfuScaler{
	hostmap: make(map[string]*Host, 0), //TODO remove host pointer, we can just using string itself maybe as we always sync host to get latest values
}

func (sfu SfuScaler) GetSessionHost(session string) *Host {
	sfu.mu.Lock()
	defer sfu.mu.Unlock()
	return sfu.hostmap[session]
}

func (sfu SfuScaler) IsScaledSession(session string) bool {
	sfu.mu.Lock()
	defer sfu.mu.Unlock()
	_, ok := sfu.hostmap[session]
	log.Infof("IsScaledSession %v - - %v", session, ok)
	return ok
}

func (sfu SfuScaler) AssignHostToSession(session string, nhost *Host, base *Host) string {
	sfu.mu.Lock()
	defer sfu.mu.Unlock()
	scaledsessionname := fmt.Sprintf("%v-scale-%v-%v", session, len(sfu.hostmap), rand.Intn(10000000))
	sfu.hostmap[scaledsessionname] = nhost
	sfu.hostmap[session] = base
	return scaledsessionname
}

func (sfu SfuScaler) SyncHost(e *etcdCoordinator) {
	sfu.mu.Lock()
	defer sfu.mu.Unlock()
	for session, host := range sfu.hostmap {
		foundSession := false
		for _, livesession := range e.sessions {
			if session == livesession.Name && host.Ip == livesession.Host && host.Port == livesession.Port {
				if livesession.PeerCount > 0 {
					foundSession = true
					break
				}
			}
		}
		if !foundSession {
			_, ok := sfu.hostmap[session]
			if ok {
				log.Infof("host removed from scaler %v %v", session, sfu.hostmap[session])
				delete(sfu.hostmap, session)
			}
		}

	}

	for _, livesession := range e.sessions {
		//TODO what happens is that hostmap has the memory address of Host form when it was assigned,
		// but if in the mean while more tracks get added to it, those don't get updated on this host
		// so we have to assign the latest host every time
		if strings.Index(livesession.Name, "-scale-") != -1 {
			log.Infof("found scaled session %v", livesession.Name)
			hostfound := false
			origin := ""
			for _, host := range e.hosts {
				if host.Ip == livesession.Host && host.Port == livesession.Port {
					if livesession.PeerCount > 0 {
						sfu.hostmap[livesession.Name] = &host
						log.Infof("assigned host to scaled session %v - %v", livesession.Name, host.String())
						hostfound = true
						break
					}
				}
			}
			if !hostfound {
				log.Infof("not host found for scaled session")
				_, ok := sfu.hostmap[livesession.Name]
				if ok {
					log.Infof("host removed from scaler %v %v", livesession.Name, sfu.hostmap[livesession.Name])
					delete(sfu.hostmap, livesession.Name)
				}
			} else {
				log.Infof("host found for scaled session")
				idx := strings.Index(livesession.Name, "-scale-")
				origin = livesession.Name[:idx]
			}

			hostfound = false
			for _, livesession := range e.sessions {
				if livesession.Name == origin {
					for _, host := range e.hosts {
						if host.Ip == livesession.Host && host.Port == livesession.Port {
							if livesession.PeerCount > 0 {
								sfu.hostmap[origin] = &host
								hostfound = true
								log.Infof("assigned host to origin session %v - %v", origin, host.String())
								break
							}
						}
					}
					break
				}
			}
			if !hostfound {
				_, ok := sfu.hostmap[origin]
				if ok {
					log.Infof("host removed from scaler %v %v", origin, sfu.hostmap[origin])
					delete(sfu.hostmap, origin)
				}
			}

		}

	}
}

func (sfu SfuScaler) GetOptimalSourceHost(session string) *Host {
	sfu.mu.Lock()
	defer sfu.mu.Unlock()
	log.Infof("finding optmizal host for %v", session)
	origin := session
	if strings.Index(session, "-scale-") != -1 {
		idx := strings.Index(session, "-scale-")
		origin = session[:idx]
	}

	log.Infof("origin %v", origin)

	allhosts := make(map[string]*Host)
	allhosts[origin] = sfu.hostmap[origin]

	log.Infof("sfu hostmap %v", sfu.hostmap)
	for name, host := range sfu.hostmap {
		log.Infof("name %v origin %v", name, origin)
		if strings.Index(name, origin+"-scale-") != -1 {
			allhosts[name] = host
		}
	}

	log.Infof("allhosts %v", allhosts)

	var min_load_host *Host
	var min_load float64
	for _, host := range allhosts {
		cpu := host.GetCurrentLoad()
		log.Infof("host %v loads %v", host.Ip, cpu)
		if cpu > min_load {
			min_load_host = host
		}
	}

	log.Infof("optmial host found for scaling", min_load_host.String())

	return min_load_host
}

func (sfu SfuScaler) FindOptmialHost(session string, e *etcdCoordinator) (*Host, string) {
	sfu.mu.Lock()
	defer sfu.mu.Unlock()
	log.Infof("finding optmizal host for %v", session)
	origin := session
	if strings.Index(session, "-scale-") != -1 {
		idx := strings.Index(session, "-scale-")
		origin = session[:idx]
	}

	log.Infof("origin %v", origin)

	allhosts := make(map[string]*Host)
	allhosts[origin] = sfu.hostmap[origin]

	log.Infof("sfu hostmap %v", sfu.hostmap)
	for name, host := range sfu.hostmap {
		log.Infof("name %v origin %v", name, origin)
		if strings.Index(name, origin+"-scale-") != -1 {
			allhosts[name] = host
		}
	}

	log.Infof("allhosts %v", allhosts)

	var optmialhost *Host
	opsession := ""
	for name, host := range allhosts {
		if e.canHostServe(host) {
			optmialhost = host
			opsession = name
			break
		}
	}

	return optmialhost, opsession
}
