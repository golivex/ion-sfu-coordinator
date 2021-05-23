package main

import (
	"context"

	coordinator "github.com/golivex/sfu-coordinator/services"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	etcd := coordinator.NewCoordinatorEtcd("0.0.0.0:2379")
	etcd.LoadSessions()
	etcd.LoadHosts()
	go etcd.WatchHosts(ctx)
	go etcd.WatchSessions(ctx)
	go etcd.WatchCurrentSessionMap(ctx)

	defer cancel()
	defer etcd.Close()
	etcd.InitApi()

}
