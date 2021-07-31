package coordinator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/pion/ion-log"
)

func (e *etcdCoordinator) getRtmpKey() string {

	usedActions := make(map[string]int)
	actionhost := e.getReadyActionHost()
	if actionhost != nil {
		_, ok := usedActions[actionhost.Ip]
		if ok {
			actionhost = nil
		}
	}
	if actionhost == nil {
		usedActions["CLOUD_START_rtmp"] = 1
		go func() {
			notifyip := e.startActionHost(-1, "rtmptotrack") //start 2vcpu machine
			log.Infof("waiting for action machine ip")
			ip := <-notifyip
			log.Infof("got action machine ip %v", ip)
			actionhost := e.getActionHostByIp(ip)
			if actionhost == nil {
				panic("host cannot be nil!")
			}
			e.getKeyOnHost(actionhost.Ip, actionhost.Port, 3)
		}()
	} else {
		log.Infof("action host found %v", actionhost.String())
		usedActions[actionhost.Ip] = 1
		key := e.getKeyOnHost(actionhost.Ip, actionhost.Port, 3)
		usedActions[key] = 1
	}
	b, _ := json.Marshal(usedActions)
	return string(b)
}

func (e *etcdCoordinator) startRtmp(session string, rtmp string) string {

	usedActions := make(map[string]int)
	actionhost := e.getReadyActionHost()
	if actionhost != nil {
		_, ok := usedActions[actionhost.Ip]
		if ok {
			actionhost = nil
		}
	}
	if actionhost == nil {
		usedActions["CLOUD_START_rtmp"] = 1
		go func() {
			notifyip := e.startActionHost(-1, "rtmptotrack") //start 2vcpu machine
			log.Infof("waiting for action machine ip")
			ip := <-notifyip
			log.Infof("got action machine ip %v", ip)
			actionhost := e.getActionHostByIp(ip)
			if actionhost == nil {
				panic("host cannot be nil!")
			}
			e.startRtmpOnHost(session, rtmp, actionhost.Ip, actionhost.Port, 3)
			//TODO there is an issue even if we get the key here we cannot return this to the api
			// so there will be an empty actions server started for rtmp but it won't be doing anything
			//its possible some other api takes this action server during high traffic,
			// but for now this should be fine
			// as from frontend we will keep hitting the api till we get a key
		}()
	} else {
		log.Infof("action host found %v", actionhost.String())
		usedActions[actionhost.Ip] = 1
		e.startRtmpOnHost(session, rtmp, actionhost.Ip, actionhost.Port, 3)
	}
	b, _ := json.Marshal(usedActions)
	return string(b)
}

func (e *etcdCoordinator) getKeyOnHost(host, port string, retry int) string {
	apiurl := "http://" + host + ":" + port + "/rtmp/getkey/zoom"
	log.Infof("rtmp api called %v retry %v", apiurl, retry)
	resp, err := http.Get(apiurl)
	if err != nil {
		log.Errorf("%v", err)
		if retry > 1 {
			time.Sleep(5) //it takes time for host to get ready
			return e.getKeyOnHost(host, port, retry-1)
		}
		return fmt.Sprintf("Err %v", err)
	}
	log.Infof("get rtmp key %v", resp.Body)
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(responseData)
}

func (e *etcdCoordinator) startRtmpOnHost(session, rtmp, host, port string, retry int) string {
	apiurl := "http://" + host + ":" + port + "/rtmp/live/" + session + "/" + rtmp
	if len(rtmp) == 0 {
		apiurl = "http://" + host + ":" + port + "/rtmp/demo/" + session
	}

	log.Infof("rtmp api called %v retry %v", apiurl, retry)
	resp, err := http.Get(apiurl)
	if err != nil {
		log.Errorf("%v", err)
		if retry > 1 {
			time.Sleep(5) //it takes time for host to get ready
			return e.startStreamOnHost(session, rtmp, host, port, retry-1)
		}
		return fmt.Sprintf("Err %v", err)
	}
	log.Infof("save to disk %v", resp.StatusCode)
	return resp.Status
}

func (e *etcdCoordinator) stopRtmp(session string) string {
	return e.stopAction(session, "rtmptotrack")
}
