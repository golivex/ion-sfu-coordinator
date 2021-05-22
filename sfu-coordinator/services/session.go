package coordinator

import (
	"context"
	"fmt"
	"strings"
	"time"

	log "github.com/pion/ion-log"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func (e *etcdCoordinator) deleteSessionsForHost(host string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	for key, _ := range e.sessions {
		if strings.Index(key, host) != -1 {
			delete(e.sessions, key)
		}
	}
}

func (e *etcdCoordinator) getKey(sessionstr string) string {
	s := strings.Split(sessionstr, "/")

	var sessionname string
	if len(s) > 2 {
		sessionname = s[2]
	}
	key := sessionname
	if len(s) > 4 {
		ip := s[4]
		key = sessionname + "-" + ip
	}
	return key

}

func (e *etcdCoordinator) removeSession(sessionstr string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	s := strings.Split(sessionstr, "/")

	key := e.getKey(sessionstr)

	fmt.Println("key", key)
	_, ok := e.sessions[key]
	if !ok {
		log.Infof("session doesnt exist %v %v", key, e.sessions)
		return
	}
	if len(s) > 6 {
		p := s[6]

		foundidx := -1
		for idx, pp := range e.sessions[key].Peers {
			if pp.Id == p {
				foundidx = idx
				break
			}
		}
		fmt.Println("ePeer", foundidx)

		if len(s) > 8 {
			if len(s) > 10 {
				t := s[8]
				kind := s[10]

				if foundidx != -1 {
					foundTrackIdx := -1
					ePeer := e.sessions[key].Peers[foundidx]
					for idx, track := range ePeer.Tracks {
						if track.Id == t && track.Kind == kind {
							foundTrackIdx = idx
						}
					}
					ePeer.Tracks = append(ePeer.Tracks[:foundTrackIdx], ePeer.Tracks[foundTrackIdx+1:]...)
					e.sessions[key].Peers[foundidx] = ePeer

					if foundTrackIdx == -1 {
						fmt.Println(sessionstr)
						panic("track not found")
					}
				} else {
					fmt.Println(sessionstr)
					panic("peer not found")
				}

			} else {
				log.Errorf("Invalid session key2 %v", sessionstr)
				panic("invalid session key2")
			}
		} else {
			if foundidx != -1 {
				fmt.Println("removing peer %v", p)
				livesession := e.sessions[key]
				livesession.Peers = append(livesession.Peers[:foundidx], livesession.Peers[foundidx+1:]...)
				e.sessions[key] = livesession
			}
		}

	} else {
		delete(e.sessions, key)
	}
	e.debugSession()
	e.updateHostSessions()
}

func (e *etcdCoordinator) generateSessionTree(sessionstr string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	s := strings.Split(sessionstr, "/")

	key := e.getKey(sessionstr)

	fmt.Println("key", key)

	var ip, port, sessionname string

	if len(s) > 2 {
		sessionname = s[2]
	}
	if len(s) > 4 {
		ip = s[4]
		port = strings.Split(ip, ":")[1]
		ip = strings.Split(ip, ":")[0]
	}

	_, ok := e.sessions[key]
	if !ok {
		e.sessions[key] = LiveSession{
			Name:        sessionname,
			Host:        ip,
			Port:        port,
			PeerCount:   0,
			AudioTracks: 0,
			VideoTracks: 0,
			Peers:       []Peer{},
		}
	}
	if len(s) > 6 {
		p := s[6]

		var ePeer *Peer
		foundidx := -1
		for idx, pp := range e.sessions[key].Peers {
			if pp.Id == p {
				ePeer = &pp
				foundidx = idx
				break
			}
		}
		fmt.Println("ePeer", ePeer)
		if ePeer == nil {
			ePeer = &Peer{
				Id:     p,
				Tracks: make([]Track, 0),
			}
		}

		if len(s) > 8 {
			if len(s) > 10 {
				t := s[8]
				kind := s[10]
				if ePeer.Tracks == nil {
					ePeer.Tracks = []Track{}
					ePeer.Tracks[0] = Track{
						Id:   t,
						Kind: kind,
					}
				} else {
					ePeer.Tracks = append(ePeer.Tracks, Track{
						Id:   t,
						Kind: kind,
					})
				}
			} else {
				panic(fmt.Sprintf("unexpected session string %v", sessionstr))
			}

		}

		livesession := e.sessions[key]

		if foundidx == -1 {
			fmt.Println("e.sessions[key].peer", livesession.Peers)
			livesession.Peers = append(livesession.Peers, *ePeer)
		} else {
			livesession.Peers[foundidx] = *ePeer
		}

		audioTracks := 0
		videoTracks := 0
		for _, peer := range livesession.Peers {
			for _, track := range peer.Tracks {
				if track.Kind == "audio" {
					audioTracks = audioTracks + 1
				}
				if track.Kind == "video" {
					videoTracks = videoTracks + 1
				}
			}

		}

		livesession.AudioTracks = audioTracks
		livesession.VideoTracks = videoTracks
		livesession.PeerCount = len(livesession.Peers)

		e.sessions[key] = livesession
	}
	e.debugSession()
	e.updateHostSessions()
}

func (e *etcdCoordinator) updateHostSessions() {
	for hostkey, host := range e.hosts {
		host.AudioTracks = 0
		host.VideoTracks = 0
		host.PeerCount = 0
		for _, live := range e.sessions {
			// log.Infof("host.Ip %v live.Host %v host.Port %v live.Port %v", host.Ip, live.Host, host.Port, live.Port)
			if host.Ip == live.Host && host.Port == live.Port {
				host.AudioTracks = host.AudioTracks + live.AudioTracks
				host.VideoTracks = host.VideoTracks + live.VideoTracks
				host.PeerCount = host.PeerCount + live.PeerCount
			}

		}
		e.hosts[hostkey] = host
	}

}

func (e *etcdCoordinator) debugSession() {
	fmt.Println("=================================")
	for _, live := range e.sessions {
		fmt.Println(fmt.Sprintf("======== session:%v host:%v port:%v ========", live.Name, live.Host, live.Port))
		fmt.Println(fmt.Sprintf("======== peer count:%v audioTracks:%v videoTracks %v ========", live.PeerCount, live.AudioTracks, live.VideoTracks))
		for _, peer := range live.Peers {
			fmt.Println(fmt.Sprintf("\t====peer-%v====", peer.Id))
			for _, track := range peer.Tracks {
				fmt.Println(fmt.Sprintf("\t\t==track:%v:%v==", track.Id, track.Kind))
			}
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func (e *etcdCoordinator) LoadSessions() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := e.cli.Get(ctx, "/session/", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()
	if err != nil {
		log.Errorf("error fetching session", err)
	}
	for _, ev := range resp.Kvs {
		// fmt.Printf("%s : %s\n", ev.Key, ev.Value)
		sessionstr := string(ev.Key[:])
		fmt.Println("session str %v", sessionstr)
		e.generateSessionTree(sessionstr)
	}
}

func (e *etcdCoordinator) WatchSessions(ctx context.Context) {
	log.Infof("watching sessions")
	rch := e.cli.Watch(ctx, "/session/", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			// log.Infof("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			sessionstr := string(ev.Kv.Key[:])
			log.Infof("%v session str %v", ev.Type, sessionstr)
			// /session/test/node/5.9.18.28:7002/peer/ckoy35usg00080110qpo13b3v
			if ev.Type == mvccpb.PUT {
				e.generateSessionTree(sessionstr)

			}
			if ev.Type == mvccpb.DELETE {
				e.removeSession(sessionstr)
			}
		}
	}
}
