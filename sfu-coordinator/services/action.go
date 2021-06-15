package coordinator

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/pion/ion-log"
)

type ActionStatus struct {
	IsActive   bool   `json:"isactive"`
	Session    string `json:"session"`
	Err        error  `json:"err"`
	ActionType string `json:"tasktype"`
	Progress   string `json:"progress"`
}

func (e *etcdCoordinator) queryActionStatus(session string, action string) (*Host, *ActionStatus) {
	for _, host := range e.actionhosts {
		actionstatus := e.getActionStatus(host.Ip, host.Port)
		if actionstatus.IsActive && actionstatus.Session == session && actionstatus.ActionType == action {
			return &host, actionstatus
		}
	}
	return nil, nil
}

func (e *etcdCoordinator) getActionStatus(host string, port string) *ActionStatus {
	apiurl := "http://" + host + ":" + port + "/status"
	log.Infof("api called %v", apiurl)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(apiurl)
	if err != nil {
		log.Errorf("%v", err)
		return nil
	}
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Errorf("%v", err)
		return nil
	} else {
		var response ActionStatus
		err = json.Unmarshal(body, &response)
		return &response
	}
}

func (e *etcdCoordinator) stopAction(session string, actionType string) string {
	for _, host := range e.actionhosts {
		status := e.getActionStatus(host.Ip, host.Port)
		if status != nil {
			log.Infof("status %v", status)
			if status.IsActive && status.ActionType == actionType {
				apiurl := "http://" + host.Ip + ":" + host.Port + "/stop"
				log.Infof("stop action api called %v", apiurl)
				client := http.Client{
					Timeout: 5 * time.Second,
				}
				client.Get(apiurl)
				return host.String()
			}
		}
	}
	return "NO_HOST_FOUND"
}
