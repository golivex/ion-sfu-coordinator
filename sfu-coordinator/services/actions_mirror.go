package coordinator

import (
	"net/http"
	"sync"
	"time"

	log "github.com/pion/ion-log"
)

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

	apiurl := "http://" + host.Ip + ":" + host.Port + "/mirror/syncsfu/" + session + "/" + session2 + "/" + host.String() + "/" + nhost.String()
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
