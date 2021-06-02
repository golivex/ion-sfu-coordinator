package main

import (
	"flag"
	"time"

	tasktodisk "github.com/manishiitg/actions/tracktodisk"
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

	tasktodisk.Init(session, addr, cancel)

	time.AfterFunc(6*time.Second, func() {
		close(cancel)
	})
	<-cancel
}
