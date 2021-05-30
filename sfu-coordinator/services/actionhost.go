package coordinator

import (
	"context"

	log "github.com/pion/ion-log"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func (e *etcdCoordinator) getReadyActionHost() *Host {
	var h *Host
	for _, host := range e.actionhosts {
		if host.Tasks == 0 {
			if e.checkActionNode(host.Ip, host.Port) {
				h = &host
			}
		}
	}
	return h
}

func (e *etcdCoordinator) getActionHostByIp(ip string) *Host {
	var h *Host
	for _, host := range e.actionhosts {
		log.Infof("host ip %v ip %v host debug", host.Ip, ip, host.String())
		if host.Ip == ip {
			h = &host
			break
		}
	}
	return h
}

func (e *etcdCoordinator) startActionHost(capacity int) chan string {
	notifyip := make(chan string, 1)
	log.Infof("startActionHost %v", capacity)
	go func() {
		log.Infof("starting action machine with capacity %v", capacity)
		if e.cloud.StartActionServerNotify(capacity, notifyip) {
			log.Infof("waiting for ip of action machine")
		} else {
			log.Infof("unable to start action machine")
		}
	}()
	return notifyip
}

func (e *etcdCoordinator) WatchActionHosts(ctx context.Context) {
	log.Infof("watching action hosts")
	rch := e.cli.Watch(ctx, "action-hosts/", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			if ev.Type == mvccpb.PUT {
				ip := string(ev.Kv.Key[:])
				loadStr := ev.Kv.Value[:]
				e.addHost(ip, loadStr, true)
			}
			if ev.Type == mvccpb.DELETE {
				ip := string(ev.Kv.Key[:])
				e.deleteHostString(ip, true)

			}
			if e.cloud != nil {
				for _, host := range e.actionhosts {
					cpu := host.GetCurrentLoad()
					e.cloud.UpdateActionNodeLoad(host.Ip, host.Port, host.Tasks, cpu)
				}
			}
			// log.Infof(" watch host %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			// log.Infof("hosts %v", e.hosts)
		}
	}
}
