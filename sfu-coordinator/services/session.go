package coordinator

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	log "github.com/pion/ion-log"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Track struct {
	Id   string
	Kind string
}

type LiveSession struct {
	Name        string
	Host        string
	Port        string
	PeerCount   int
	AudioTracks int
	VideoTracks int
	Peers       []Peer
}

type Peer struct {
	Id     string
	Tracks []Track
}

func (e *etcdCoordinator) deleteOrphanSession() {
	e.mu.Lock()

	log.Infof("deleteOrphanSession")
	for key, session := range e.sessions {
		exist := false
		for _, host := range e.hosts {
			if host.Ip == session.Host && host.Port == session.Port {
				exist = true
				break
			}
		}
		if !exist {
			log.Infof("orphan session found %v", session.Name)
			delete(e.sessions, key)
		}
	}
	e.mu.Unlock()
}

func (e *etcdCoordinator) deleteSessionsForHost(host string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	log.Infof("deleteSessionsForHost %v", host)
	for key, _ := range e.sessions {
		if strings.Index(key, host) != -1 {
			delete(e.sessions, key)
		}
	}
	e.updateHostSessions()
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

	log.Infof("key %v", key)
	_, ok := e.sessions[key]
	if !ok {
		// log.Infof("session doesnt exist %v %v", key, e.sessions)
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
		log.Infof("ePeer %v", foundidx)

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
						log.Errorf(sessionstr)
						panic("track not found")
					}
				} else {
					log.Errorf(sessionstr)
					// panic("peer not found")
				}

			} else {
				log.Errorf("Invalid session key2 %v", sessionstr)
				panic("invalid session key2")
			}
		} else {
			if foundidx != -1 {
				log.Infof("removing peer %v", p)
				livesession := e.sessions[key]
				livesession.Peers = append(livesession.Peers[:foundidx], livesession.Peers[foundidx+1:]...)

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
		}

	} else {
		delete(e.sessions, key)
	}
	e.updateHostSessions()

	// e.debugSession()
}

func (e *etcdCoordinator) generateSessionTree(sessionstr string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	s := strings.Split(sessionstr, "/")

	key := e.getKey(sessionstr)

	// log.Infof("key %v", key)

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
		log.Infof("ePeer %v", ePeer)
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
				log.Errorf("unexpected session string %v", sessionstr)
				panic("unexpected session string")
			}

		}

		livesession := e.sessions[key]

		if foundidx == -1 {
			// log.Infof("e.sessions[key].peer %v", livesession.Peers)
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
	e.updateHostSessions()
	// e.debugSession()
}

func (e *etcdCoordinator) updateSessionMap(stats map[string][]string, node string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	for sessionkey, live := range e.sessions {
		if live.Host+":"+live.Port == node {
			// log.Infof("looking at node %v", node)

			sessionfound := false
			for key, _ := range stats {
				if live.Name == key {
					sessionfound = true
					// log.Infof("%v session found in host looking for peers", sessionkey)
					var newPeers = []Peer{}
					for _, peer := range live.Peers {
						peerfound := false
						for _, rp := range stats[key] {
							if rp == peer.Id {
								peerfound = true
								break
							}
						}
						if peerfound {
							newPeers = append(newPeers, peer)
						} else {
							log.Infof("removing peer! %v", peer.Id)
						}
					}
					live.Peers = newPeers

					audioTracks := 0
					videoTracks := 0
					for _, peer := range live.Peers {
						for _, track := range peer.Tracks {
							if track.Kind == "audio" {
								audioTracks = audioTracks + 1
							}
							if track.Kind == "video" {
								videoTracks = videoTracks + 1
							}
						}

					}

					live.AudioTracks = audioTracks
					live.VideoTracks = videoTracks
					live.PeerCount = len(live.Peers)

					e.sessions[sessionkey] = live
					break
				}
			}

			if !sessionfound {
				// log.Infof("%v session not found in host, removing it!!!!", sessionkey)
				delete(e.sessions, sessionkey)
			}

		}
	}

	e.updateHostSessions()
	// e.debugSession()
}

func (e *etcdCoordinator) updateHostSessions() {
	for hostkey, host := range e.hosts {
		host.AudioTracks = 0
		host.VideoTracks = 0
		host.PeerCount = 0
		for _, live := range e.sessions {
			if host.Ip == live.Host && host.Port == live.Port {
				// log.Infof("host.Ip %v live.Host %v host.Port %v live.Port %v live %v", host.Ip, live.Host, host.Port, live.Port, live)
				host.AudioTracks = host.AudioTracks + live.AudioTracks
				host.VideoTracks = host.VideoTracks + live.VideoTracks
				host.PeerCount = host.PeerCount + live.PeerCount
			}
		}
		e.hosts[hostkey] = host
	}

	if e.cloud != nil {
		for _, host := range e.hosts {
			cpu := host.GetCurrentLoad()
			e.cloud.UpdateNodeLoad(host.Ip, host.Port, host.PeerCount, cpu)
		}
	}

}

func (e *etcdCoordinator) debugSession() {
	log.Infof("==============SESSIONS===================")
	for _, live := range e.sessions {
		log.Infof("======== session:%v host:%v port:%v ========", live.Name, live.Host, live.Port)
		log.Infof("======== peer count:%v audioTracks:%v videoTracks %v ========", live.PeerCount, live.AudioTracks, live.VideoTracks)
		for _, peer := range live.Peers {
			log.Infof("\t====peer-%v====", peer.Id)
			for _, track := range peer.Tracks {
				log.Infof("\t\t==track:%v:%v==", track.Id, track.Kind)
			}
		}
	}

	log.Infof("=================================")
	log.Infof("==============HOSTS===================")

	for _, host := range e.hosts {
		log.Infof("======== host:%v ========", host.String())
		log.Infof("======== peer %v atracks:%v vtracks:%v ========", host.PeerCount, host.AudioTracks, host.VideoTracks)
	}
	log.Infof("*************")
}

func (e *etcdCoordinator) LoadSessions() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := e.cli.Get(ctx, "/session/", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()
	if err != nil {
		log.Errorf("error fetching session", err)
	}
	for _, ev := range resp.Kvs {
		// log.Infof("%s : %s\n", ev.Key, ev.Value)
		sessionstr := string(ev.Key[:])
		log.Infof("load session str %v", sessionstr)
		e.generateSessionTree(sessionstr)
	}
	e.deleteOrphanSession()
	ticker := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-ticker.C:
			e.deleteOrphanSession()
		}
	}
}

func (e *etcdCoordinator) WatchCurrentSessionMap(ctx context.Context) {
	log.Infof("watching session map")
	rch := e.cli.Watch(ctx, "/current_session_map/", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			// log.Infof("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			// /session/test/node/5.9.18.28:7002/peer/ckoy35usg00080110qpo13b3v
			sessionstr := string(ev.Kv.Key[:])
			var stats = make(map[string][]string)
			json.Unmarshal(ev.Kv.Value, &stats)
			node := strings.Split(sessionstr, "/")[3]
			// log.Infof("%v current_session_map str %v %v %v", ev.Type, sessionstr, stats, node)
			if ev.Type == mvccpb.PUT {
				e.updateSessionMap(stats, node)
			}
			if ev.Type == mvccpb.DELETE {
			}
		}
	}
}

func (e *etcdCoordinator) WatchSessions(ctx context.Context) {
	log.Infof("watching sessions")
	rch := e.cli.Watch(ctx, "/session/", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			// log.Infof("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			sessionstr := string(ev.Kv.Key[:])
			// log.Infof("%v session str %v", ev.Type, sessionstr)
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
