package main

import (
	"flag"

	loadtest "github.com/manishiitg/actions/loadtest"
	"github.com/manishiitg/actions/loadtest/client/gst"
	log "github.com/pion/ion-log"
)

func clientThread() {
	//get args
	var session, gaddr, file, role, loglevel, simulcast, paddr string
	var total, cycle, duration int
	var video, audio bool

	var create_room = -1

	flag.StringVar(&file, "file", "test", "Path to the file media")
	flag.StringVar(&gaddr, "gaddr", "http://0.0.0.0:4000/", "Ion-sfu grpc addr")
	flag.StringVar(&session, "session", "test", "join session name")
	flag.IntVar(&total, "clients", 1, "Number of clients to start")
	flag.IntVar(&cycle, "cycle", 1000, "Run new client cycle in ms")
	flag.IntVar(&duration, "duration", 60*60, "Running duration in sencond")
	flag.StringVar(&role, "role", "pubsub", "Run as pubsub/sub")
	flag.StringVar(&loglevel, "log", "info", "Log level")
	flag.BoolVar(&video, "v", true, "Publish video stream from webm file")
	flag.BoolVar(&audio, "a", true, "Publish audio stream from webm file")
	flag.StringVar(&simulcast, "simulcast", "", "simulcast layer q|h|f")
	flag.StringVar(&paddr, "paddr", "", "pprof listening addr")
	flag.IntVar(&create_room, "create_room", -1, "number of peers per room")
	flag.Parse()
	log.Init(loglevel)

	cancel := make(chan struct{})
	go loadtest.Init(file, gaddr, session, total, cycle, duration, role, video, audio, simulcast, paddr, create_room, cancel)
	// Listen for signals

}

func main() {
	go clientThread()
	gst.MainLoop()
}
