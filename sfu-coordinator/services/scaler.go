package coordinator

import (
	"fmt"
	"math/rand"
	"strings"

	log "github.com/pion/ion-log"
)

type SfuScaler struct {
	hostmap map[string]Host //this will map all scaled session and their assigned hosts
}

var sfu = SfuScaler{
	hostmap: make(map[string]Host, 0),
}

func (sfu SfuScaler) GetSessionHost(session string) Host {
	return sfu.hostmap[session]
}

func (sfu SfuScaler) IsScaledSession(session string) bool {
	// return strings.Index(session, "-scale-") != -1
	_, ok := sfu.hostmap[session]
	return ok
}

func (sfu SfuScaler) AssignHostToSession(session string, nhost Host, base Host) string {
	// _, ok := sfu.origin[session]
	// if !ok {
	// 	sfu.origin[session] = []HostSession{}
	// 	scaledsessionname = fmt.Sprintf("%v-scale-%v", session, 0)
	// } else {
	// 	scaledsessionname = fmt.Sprintf("%v-scale-%v", session, len(sfu.origin[session]))
	// }
	// sfu.origin[session] = append(sfu.origin[session], HostSession{
	// 	ip:   nhost.Ip,
	// 	name: scaledsessionname,
	// })
	scaledsessionname := fmt.Sprintf("%v-scale-%v-%v", session, len(sfu.hostmap), rand.Intn(10000000))
	sfu.hostmap[scaledsessionname] = nhost
	sfu.hostmap[session] = base
	return scaledsessionname
}

func getHostFromSessionName() {

}

func (sfu SfuScaler) SyncHost(e *etcdCoordinator) {
	for session, host := range sfu.hostmap {
		foundSession := false
		for _, livesession := range e.sessions {
			if session == livesession.Name && host.Ip == livesession.Host && host.Port == livesession.Port {
				foundSession = true
				break
			}
		}
		if !foundSession {
			log.Infof("host removed from scaler %v %v", session, sfu.hostmap[session])
			delete(sfu.hostmap, session)
		}

	}

	for _, livesession := range e.sessions {

		if strings.Index(livesession.Name, "-scale-") != -1 {
			log.Infof("found scaled session %v", livesession.Name)
			_, ok := sfu.hostmap[livesession.Name]
			if !ok {
				for _, host := range e.hosts {
					if host.Ip == livesession.Host && host.Port == livesession.Port {
						sfu.hostmap[livesession.Name] = host
						log.Infof("assigned host to scaled session %v - %v", livesession.Name, host.String())
						break
					}
				}
			}

			idx := strings.Index(livesession.Name, "-scale-")
			origin := livesession.Name[:idx]
			_, ok = sfu.hostmap[origin]

			if !ok {
				for _, livesession := range e.sessions {
					if livesession.Name == origin {
						for _, host := range e.hosts {
							if host.Ip == livesession.Host && host.Port == livesession.Port {
								sfu.hostmap[origin] = host
								log.Infof("assigned host to origin session %v - %v", origin, host.String())
								break
							}
						}
						break
					}
				}
			}

		}

	}
}

func (sfu SfuScaler) FindOptmialHost(session string, e *etcdCoordinator) Host {
	log.Infof("finding optmizal host for %v", session)
	origin := session
	if strings.Index(session, "-scale-") != -1 {
		idx := strings.Index(session, "-scale-")
		origin = session[:idx]
	}

	log.Infof("origin %v", origin)

	allhosts := []Host{}
	allhosts = append(allhosts, sfu.hostmap[origin])

	log.Infof("%v", sfu.hostmap)
	for name, host := range sfu.hostmap {
		log.Infof("name %v origin %v", name, origin)
		if strings.Index(name, origin+"-scale-") != -1 {
			allhosts = append(allhosts, host)
		}
	}

	log.Infof("allhosts %v", allhosts)

	var optmialhost Host

	for _, host := range allhosts {
		if e.canHostServe(host) {
			optmialhost = host
			break
		}
	}

	return optmialhost
}
