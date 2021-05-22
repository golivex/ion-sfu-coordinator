package main

import (
	"flag"

	"github.com/manishiitg/actions/mirror-sfu/client"
	log "github.com/pion/ion-log"
)

func main() {
	// init log
	log.Init("info")

	// parse flag
	var session, session2, addr, addr2 string
	flag.StringVar(&addr, "addr", "5.9.18.28:50052", "Ion-sfu grpc addr")
	flag.StringVar(&session, "session", "test", "join session name")

	flag.StringVar(&addr2, "addr2", "5.9.18.28:50052", "Ion-sfu grpc addr")
	flag.StringVar(&session2, "session2", "test2", "join session name")
	flag.Parse()

	client.InitWithAddress(session, session2, addr, addr2)
}
