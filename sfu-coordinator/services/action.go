package coordinator

import (
	"net/http"

	log "github.com/pion/ion-log"
)

const ACTION_PORT = ":3050"

func MirrorSfu(session, session2 string, host Host, nhost Host) {
	resp, err := http.Get("http://" + host.Ip + ACTION_PORT + "/syncsfu/" + session + "/" + session2 + "/" + host.String() + "/" + nhost.String())
	if err != nil {
		log.Errorf("%v", err)
		panic(err)
	}
	log.Infof("mirror sfu %v", resp.StatusCode)
}

func SimLoad(session string, host string) {
	resp, err := http.Get("http://" + host + ACTION_PORT + "/load/" + session)
	if err != nil {
		log.Errorf("%v", err)
		panic(err)
	}
	log.Infof("SimLoad sfu %v", resp.StatusCode)
}

func StopSimLoad(host string) {
	resp, err := http.Get("http://" + host + ACTION_PORT + "/stopload")
	if err != nil {
		log.Errorf("%v", err)
		panic(err)
	}
	log.Infof("SimLoad sfu %v", resp.StatusCode)
}
