package main

import (
	"flag"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/manishiitg/actions/loadtest/client/gst"
	tasktodisk "github.com/manishiitg/actions/tracktodisk"
	log "github.com/pion/ion-log"
)

func compositeThread(session string, addr string) {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	cancel := make(chan struct{})
	tasktodisk.Init(session, addr, "gstreamer", cancel)

	for {
		select {
		case sig := <-sigs:
			log.Infof("got signal %v", sig)
			close(cancel)
		}
	}

}

func main() {
	// init log
	log.Init("info")

	// parse flag
	var session, addr string
	flag.StringVar(&addr, "addr", "http://0.0.0.0:4000/", "SFU Cordinator")
	flag.StringVar(&session, "session", "test2", "join session name")

	flag.Parse()

	runtime.LockOSThread()
	go compositeThread(session, addr)
	gst.MainLoop()

}
