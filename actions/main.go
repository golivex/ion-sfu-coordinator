package main

import (
	"flag"
	"os"

	actions "github.com/manishiitg/actions/pkg"
	log "github.com/pion/ion-log"
)

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return ""
}

func main() {

	var eaddr, ipaddr, port string

	flag.StringVar(&eaddr, "eaddr", "0.0.0.0:2379", "etcd ip")
	flag.StringVar(&ipaddr, "ipaddr", "5.9.18.28", "current server ip")
	flag.StringVar(&port, "port", ":3050", "API Port")
	flag.Parse()

	if eaddr == "" {
		eaddr = getEnv("eaddr")
	}
	if ipaddr == "" {
		ipaddr = getEnv("ipaddr")
	}
	if port == "" {
		port = getEnv("port")
	}

	if len(eaddr) == 0 || len(ipaddr) == 0 || len(port) == 0 {
		log.Infof("ipaddr %v, eaddr %v, port %v all requried", ipaddr, eaddr, port)
		os.Exit(-1)
	}

	log.Init("info")

	e := actions.InitEtcd(eaddr, ipaddr, port)
	e.InitApi(port)
}
