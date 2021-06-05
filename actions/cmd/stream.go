package cmd

import (
	"flag"

	tracktortp "github.com/manishiitg/actions/tracktortp"
	log "github.com/pion/ion-log"
)

func loadstream() {
	// init log
	log.Init("info")

	// parse flag
	var session, addr, rtmp string
	flag.StringVar(&addr, "addr", "http://0.0.0.0:4000/", "SFU Cordinator")
	flag.StringVar(&session, "session", "test2", "join session name")
	flag.StringVar(&rtmp, "rtmp", "rtmp://bom01.contribute.live-video.net/app/live_666332364_5791UvimKkDZW8edq8DAi4011wc4cR", "rtmp url")

	flag.Parse()
	cancel := make(chan struct{})

	tracktortp.Init(session, addr, rtmp, cancel)
}
