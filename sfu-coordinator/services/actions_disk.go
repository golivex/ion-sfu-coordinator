package coordinator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/pion/ion-log"
)

func (e *etcdCoordinator) saveSessionToDisk(session string, filename string) string {

	usedActions := make(map[string]int)
	actionhost := e.getReadyActionHost()
	if actionhost != nil {
		_, ok := usedActions[actionhost.Ip]
		if ok {
			actionhost = nil
		}
	}
	if actionhost == nil {
		usedActions["CLOUD_START_disk"] = 1
		go func() {
			notifyip := e.startActionHost(-1, "tracktodisk") //start 2vcpu machine
			log.Infof("waiting for action machine ip")
			ip := <-notifyip
			log.Infof("got action machine ip %v", ip)
			actionhost := e.getActionHostByIp(ip)
			if actionhost == nil {
				panic("host cannot be nil!")
			}
			e.saveSessionToDiskOnHost(session, filename, actionhost.Ip, actionhost.Port, 3)
		}()
	} else {
		log.Infof("action host found %v", actionhost.String())
		usedActions[actionhost.Ip] = 1
		e.saveSessionToDiskOnHost(session, filename, actionhost.Ip, actionhost.Port, 3)
	}
	b, _ := json.Marshal(usedActions)
	return string(b)
}

func (e *etcdCoordinator) saveSessionToDiskOnHost(session, filename, host, port string, retry int) string {
	apiurl := "http://" + host + ":" + port + "/disk/" + session + "?filename=" + filename
	log.Infof("disk api called %v retry %v", apiurl, retry)
	resp, err := http.Get(apiurl)
	if err != nil {
		log.Errorf("%v", err)
		if retry > 1 {
			time.Sleep(5) //it takes time for host to get ready
			return e.saveSessionToDiskOnHost(session, filename, host, port, retry-1)

		}
		return fmt.Sprintf("Err %v", err)
	}
	log.Infof("save to disk %v", resp.StatusCode)
	return resp.Status
}

func (e *etcdCoordinator) stopSessionToDisk(session string) string {
	return e.stopAction(session, "tracktodisk")
}
