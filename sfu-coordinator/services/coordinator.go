package services

import (
	"context"
	"fmt"
	"sync"
	"time"

	log "github.com/pion/ion-log"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

type etcdCoordinator struct {
	mu  sync.Mutex
	cli *clientv3.Client
}

func NewCoordinatorEtcd(host string) *etcdCoordinator {
	log.Infof("creating etcd client")
	cli, err := clientv3.New(clientv3.Config{
		DialTimeout: time.Second * 3,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
		Endpoints:   []string{host}, //TODO move to config
	})
	if err != nil {
		panic(err)
	}

	log.Infof("created etcdCoordinator")
	return &etcdCoordinator{
		cli: cli,
	}
}

func (e *etcdCoordinator) WatchHosts() {
	rch := e.cli.Watch(context.Background(), "available-hosts/", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			if ev.Type == "PUT" {

			}
			if ev.Type == "DELETE" {

			}
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

func (e *etcdCoordinator) Close() {
	e.cli.Close()
}
