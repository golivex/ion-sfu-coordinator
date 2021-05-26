package main

import (
	"context"

	cloud "github.com/golivex/sfu-coordinator/cloud"
	coordinator "github.com/golivex/sfu-coordinator/services"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	hub := cloud.RegisterHub(ctx)

	etcd := coordinator.NewCoordinatorEtcd("0.0.0.0:2379")
	defer etcd.Close()
	go etcd.LoadSessions()
	go etcd.LoadHosts()
	go etcd.WatchHosts(ctx)
	go etcd.WatchSessions(ctx)
	go etcd.WatchCurrentSessionMap(ctx)

	etcd.RegisterCloudProvider(hub)

	etcd.InitApi()

}
