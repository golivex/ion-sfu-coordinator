package coordinator

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	log "github.com/pion/ion-log"
)

const ACTION_PORT = ":3050"

var mirrorChecker = map[string]time.Time{}
var checkmu sync.Mutex

const MIRROR_RETRY_WAIT = 15

//TODO check if this is even working the service at port 3050
func MirrorSfu(session, session2 string, host Host, nhost Host) {
	checkmu.Lock()
	defer checkmu.Unlock()

	key := host.String() + ":" + session + ":" + session2
	_, ok := mirrorChecker[key]
	if ok {
		if time.Now().Sub(mirrorChecker[key]) > (MIRROR_RETRY_WAIT * time.Second) {
			delete(mirrorChecker, key)
		} else {
			log.Infof("skipping mirror as operation called recently!")
			return
		}
	}

	apiurl := "http://" + host.Ip + ACTION_PORT + "/syncsfu/" + session + "/" + session2 + "/" + host.String() + "/" + nhost.String()
	log.Infof("api called %v", apiurl)
	resp, err := http.Get(apiurl)
	if err != nil {
		log.Errorf("%v", err)
		panic(err)
	} else {
		mirrorChecker[key] = time.Now()
		log.Infof("mirror sfu %v", resp.StatusCode)
	}
}

func SimLoad(session string, host string, clients int) {
	s := strconv.Itoa(clients)
	apiurl := "http://" + host + ACTION_PORT + "/load/" + session + "?clients=" + s
	log.Infof("api called %v", apiurl)
	resp, err := http.Get(apiurl)
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
