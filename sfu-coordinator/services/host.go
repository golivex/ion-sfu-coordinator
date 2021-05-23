package coordinator

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	log "github.com/pion/ion-log"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

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
	lastload := loads[len(loads)-1]
	return lastload.Cpu
}

type HostPing struct {
	Ip   string  `json:"ip"`
	Port string  `json:"port"`
	Cpu  float64 `json:"cpu"`
	Mem  float64 `json:"mem"`
}

func (e *etcdCoordinator) addHost(key string, loadStr []byte) {
	e.mu.Lock()
	defer e.mu.Unlock()
	var hostping HostPing
	json.Unmarshal(loadStr, &hostping)

	hostping.Port = strings.Replace(hostping.Port, ":", "", -1)

	l := Load{
		Cpu: hostping.Cpu,
		Mem: hostping.Mem,
	}

	_, ok := e.hosts[key]
	if !ok {
		host := Host{
			Ip:    hostping.Ip,
			Port:  hostping.Port,
			Loads: []Load{},
			Spike: []Spike{},
		}
		host.Loads = append(host.Loads, l)
		e.hosts[key] = host
		e.updateHostSessions()
	} else {
		host := e.hosts[key]
		host.Loads = append(host.Loads, l)
		len := len(host.Loads)
		if len >= 5 {
			host.Loads = host.Loads[1:]
		}
		e.hosts[key] = host
	}
}
func (e *etcdCoordinator) deleteHost(ip string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	ip = strings.Replace(ip, "available-hosts/", "", -1)
	_, ok := e.hosts[ip]
	if ok {
		delete(e.hosts, ip)
	}
	e.deleteSessionsForHost(ip)
}

func (e *etcdCoordinator) LoadHosts() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := e.cli.Get(ctx, "available-hosts//", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()
	if err != nil {
		log.Errorf("error fetching hosts", err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
		ip := string(ev.Key[:])
		loadStr := ev.Value[:]
		e.addHost(ip, loadStr)
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
				e.addHost(ip, loadStr)
			}
			if ev.Type == mvccpb.DELETE {
				ip := string(ev.Kv.Key[:])
				e.deleteHost(ip)

			}
			// log.Infof(" watch host %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			// log.Infof("hosts %v", e.hosts)
		}
	}
}
