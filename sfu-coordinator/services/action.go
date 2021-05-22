package coordinator

import (
	"net/http"

	log "github.com/pion/ion-log"
)

const ACTION_PORT = ":3050"

func MirrorSfu(session, session2, host string) {
	resp, err := http.Get("http://" + host + ACTION_PORT + "/syncsfu/" + session + "/" + session2)
	if err != nil {
		log.Errorf("%v", err)
		panic(err)
	}
	log.Infof("mirror sfu %v", resp.StatusCode)
}
