package main

import (
	services "github.com/golivex/sfu-coordinator/services"
)

func main() {
	etcd := services.NewCoordinatorEtcd("0.0.0.0:2379")
	defer etcd.Close()
	etcd.WatchHosts()
	select {}
}
