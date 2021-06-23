package coordinator

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/pion/ion-log"
)

var mirrorChecker = map[string]time.Time{}

const MIRROR_RETRY_WAIT = 15

//mirror will work different than other actions
//because we need to not just start action to fwd but we need a sfu also
// so this will start a server with both sfu and actions
//so this is only testing as of now with an existing host

func (e *etcdCoordinator) mirrorSfu(session, session2 string) string {
	actionhost := e.getReadyActionHost()
	if actionhost == nil {
		return "No ready action host"
	}

	apiurl := "http://" + actionhost.Ip + ":" + actionhost.Port + "/mirror/sync/" + session + "/" + session2
	log.Infof("api called %v", apiurl)
	_, err := http.Get(apiurl)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	return "Started"
}

func (e *etcdCoordinator) stopMirror(session string) string {
	return e.stopAction(session, "mirrorsfu")
}

func (e *etcdCoordinator) stopMirrorOnHost(session string, host string) string {
	return e.stopActionOnHost(session, host, "mirrorsfu")
}
