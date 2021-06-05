package main

import (
	"flag"

	rtmptotrack "github.com/manishiitg/actions/rtmptotrack"
	log "github.com/pion/ion-log"
)

func main() {
	// init log
	log.Init("info")

	// parse flag
	var session, addr string
	flag.StringVar(&addr, "addr", "http://0.0.0.0:4000/", "SFU Cordinator")
	flag.StringVar(&session, "session", "test2", "join session name")

	flag.Parse()
	cancel := make(chan struct{})

	rtmptotrack.Init(session, addr, cancel)

	select {}
}
