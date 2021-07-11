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

type capacity struct {
	Cap  int       `json:"cap"`
	Time time.Time `json:"time"`
}

type Host struct {
	Ip              string  `json:"ip"`
	Port            string  `json:"port"`
	Domain          string  `json:"domain"`
	Tasks           int     `json:"tasks"`
	PeerCount       int     `json:"peer"`
	AudioTracks     int     `json:"audio"`
	VideoTracks     int     `json:"video"`
	Spike           []Spike `json:"spike"`
	BlockedCapacity map[string]capacity
	Loads           []Load
	lastPing        time.Time
}

type Load struct {
	Cpu float64 `json:"cpu"`
	Mem float64 `json:"mem"`
}

func (h *Host) String() string {
	return h.Ip + ":" + h.Port
}
func (h *Host) Empty() bool {
	return h.Ip == ""
}

func (h *Host) GetCurrentLoad() float64 {
	loads := h.Loads
	if len(loads) == 0 {
		return float64(100)
	}
	lastload := float64(0)
	//max taking worst condition
	if len(loads) > 10 {
		for _, val := range loads[len(loads)-10:] {
			if val.Cpu > lastload {
				lastload = val.Cpu
			}
		}
	} else {
		for _, val := range loads {
			if val.Cpu > lastload {
				lastload = val.Cpu
			}
		}
	}

	// avg
	// lastload := float64(100)
	// if len(loads) > 5 {
	// 	for _, val := range loads[len(loads)-5:] {
	// 		lastload = lastload + val.Cpu
	// 	}
	// 	lastload = lastload / float64(len(loads[len(loads)-5:]))
	// } else {
	// 	for _, val := range loads {
	// 		lastload = lastload + val.Cpu
	// 	}
	// 	lastload = lastload / float64(len(loads))
	// }
	return lastload
}

type HostPing struct {
	Ip     string  `json:"ip"`
	Port   string  `json:"port"`
	Domain string  `json:"domain"`
	Cpu    float64 `json:"cpu"`
	Mem    float64 `json:"mem"`
	Tasks  int     `json:"tasks"`
}

func (e *etcdCoordinator) addHost(key string, loadStr []byte, isaction bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	var hostping HostPing
	json.Unmarshal(loadStr, &hostping)

	hostping.Port = strings.Replace(hostping.Port, ":", "", -1)

	l := Load{
		Cpu: hostping.Cpu,
		Mem: hostping.Mem,
	}

	var ok bool
	if isaction {
		_, ok = e.actionhosts[key]
	} else {
		_, ok = e.hosts[key]
	}
	if !ok {
		if len(hostping.Ip) == 0 {
			log.Infof("hostping %v", hostping)
			panic("empty ip") //temp
		}
		if len(hostping.Port) == 0 {
			log.Infof("hostping %v", hostping)
			panic("empty port") //temp
		}

		host := Host{
			Ip:              hostping.Ip,
			Port:            hostping.Port,
			Tasks:           hostping.Tasks,
			Domain:          hostping.Domain,
			Loads:           []Load{},
			Spike:           []Spike{},
			BlockedCapacity: make(map[string]capacity),
			lastPing:        time.Now(),
		}
		host.Loads = append(host.Loads, l)
		if isaction {
			e.actionhosts[key] = host
		} else {
			e.hosts[key] = host
		}
	} else {
		var host Host
		if isaction {
			host = e.actionhosts[key]
		} else {
			host = e.hosts[key]
		}
		host.Tasks = hostping.Tasks
		host.Loads = append(host.Loads, l)
		host.lastPing = time.Now()
		len := len(host.Loads)
		if len >= 30 {
			host.Loads = host.Loads[1:]
		}
		if isaction {
			e.actionhosts[key] = host
		} else {
			e.hosts[key] = host
		}
	}
	e.updateHostSessions()
}

func (e *etcdCoordinator) deleteHost(h *Host, isaction bool) {
	e.mu.Lock()
	var hosts map[string]Host
	if isaction {
		hosts = e.actionhosts
	} else {
		hosts = e.hosts
	}
	for key, host := range hosts {
		if host.Ip == h.Ip && host.Port == h.Port {
			log.Infof("deleting host %v", host.String())
			if !isaction {
				e.deleteSessionsForHost(&host)
				if e.cloud != nil {
					e.cloud.DeleteNode(host.Ip, host.Port)
				}
				delete(e.hosts, key)
			} else {
				if e.cloud != nil {
					e.cloud.DeleteNode(host.Ip, host.Port)
				}
				delete(e.actionhosts, key)
			}

		}
	}
	e.mu.Unlock()
}
func (e *etcdCoordinator) deleteHostString(ip string, isaction bool) {
	ip = strings.Replace(ip, "available-hosts/", "", -1)
	if isaction {
		ip = strings.Replace(ip, "action-hosts/", "", -1)
	}
	port := ""
	if strings.Contains(ip, ":") {
		port = strings.Split(ip, ":")[1]
		ip = strings.Split(ip, ":")[0]
	}

	var hosts map[string]Host
	if isaction {
		hosts = e.actionhosts
	} else {
		hosts = e.hosts
	}

	for _, host := range hosts {
		if host.Ip == ip && host.Port == port {
			log.Infof("deleting host %v type isaction %v", host.String(), isaction)
			e.deleteHost(&host, isaction)
			break
		}
	}
}

func (e *etcdCoordinator) deleteOrphanHosts() {
	for _, host := range e.hosts {
		if time.Since(host.lastPing) > 30*time.Second {
			log.Infof("no ping from hosts since last 30 sec deleting it %v", host.String())
			e.deleteHost(&host, false)
		}
	}
}

func (e *etcdCoordinator) LoadHosts(ctx context.Context) {
	log.Infof("load existing hosts")
	resp, err := e.cli.Get(ctx, "available-hosts/", clientv3.WithPrefix())
	if err != nil {
		log.Errorf("error fetching hosts", err)
	}
	for _, ev := range resp.Kvs {
		log.Infof("%s : %s\n", ev.Key, ev.Value)
		ip := string(ev.Key[:])
		loadStr := ev.Value[:]
		e.addHost(ip, loadStr, false)
	}

	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			e.deleteOrphanHosts()
		}
	}
}

func (e *etcdCoordinator) WatchHosts(ctx context.Context) {
	log.Infof("watching hosts")
	rch := e.cli.Watch(ctx, "available-hosts/", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			if ev.Type == mvccpb.PUT {
				ip := string(ev.Kv.Key[:])
				loadStr := ev.Kv.Value[:]
				e.addHost(ip, loadStr, false)
			}
			if ev.Type == mvccpb.DELETE {
				ip := string(ev.Kv.Key[:])
				e.deleteHostString(ip, false)

			}
			// log.Infof(" watch host %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			// log.Infof("hosts %v", e.hosts)
		}
	}
}
