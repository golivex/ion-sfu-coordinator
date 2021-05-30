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

	var eaddr, ipaddr, port, saddr string

	flag.StringVar(&eaddr, "eaddr", "0.0.0.0:2379", "etcd ip")
	flag.StringVar(&ipaddr, "ipaddr", "5.9.18.28", "current server ip")
	flag.StringVar(&saddr, "saddr", "http://5.9.18.28:4000/", "server cluster ip")
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
	if saddr == "" {
		saddr = getEnv("saddr")
	}

	if len(eaddr) == 0 || len(ipaddr) == 0 || len(port) == 0 || len(saddr) == 0 {
		log.Infof("ipaddr %v, eaddr %v, port %v , saddr %v all requried", ipaddr, eaddr, port, saddr)
		os.Exit(-1)
	}

	log.Init("info")

	e := actions.InitEtcd(eaddr, ipaddr, port, saddr)
	e.InitApi(port)
}
