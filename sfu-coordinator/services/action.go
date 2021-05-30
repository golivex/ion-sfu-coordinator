package coordinator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"

	log "github.com/pion/ion-log"
)

var mirrorChecker = map[string]time.Time{}
var checkmu sync.Mutex

func (e *etcdCoordinator) checkActionNode(host string, port string) bool {
	apiurl := "http://" + host + ":" + port
	log.Infof("api called %v", apiurl)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	_, err := client.Get(apiurl)
	if err != nil {
		log.Errorf("%v", err)
		return false
	}
	return true
}

func (e *etcdCoordinator) simLoad(session string, clients int, role string, cycle int, rooms int, file string) string {
	actionhost := e.getReadyActionHost()
	if actionhost == nil {
		go func() {
			notifyip := e.startActionHost(-1)
			log.Infof("waiting for action machine ip")
			ip := <-notifyip
			log.Infof("got action machine ip %v", ip)
			actionhost := e.getActionHostByIp(ip)
			if actionhost == nil {
				panic("host cannot be nil!")
			}
			e.simLoadForHost(session, actionhost.Ip, actionhost.Port, clients, role, cycle, rooms, file)
		}()
		return "NEW_CLOUD_HOST_STARTED"
	} else {
		e.simLoadForHost(session, actionhost.Ip, actionhost.Port, clients, role, cycle, rooms, file)
		return actionhost.String()
	}
}

func (e *etcdCoordinator) simLoadForHost(session string, host string, port string, clients int, role string, cycle int, rooms int, file string) string {
	apiurl := "http://" + host + ":" + port + "/load/" + session + "?clients=" + strconv.Itoa(clients) + "&role=" + role + "&cycle=" + strconv.Itoa(cycle) + "&rooms=" + strconv.Itoa(rooms) + "&file=" + file
	log.Infof("api called %v", apiurl)
	resp, err := http.Get(apiurl)
	if err != nil {
		log.Errorf("%v", err)
		return fmt.Sprintf("Err %v", err)
	}
	log.Infof("SimLoad sfu %v", resp.StatusCode)
	return resp.Status
}

func (e *etcdCoordinator) stopAllSimLoad() []string {
	var stopped []string
	for _, h := range e.actionhosts {
		stopped = append(stopped, h.Ip+":"+h.Port)
		go e.stopSimLoad(h.Ip, h.Port)
	}
	return stopped
}

func (e *etcdCoordinator) stopSimLoad(host string, port string) string {
	found := false
	for _, h := range e.actionhosts {
		if h.Ip == host && h.Port == port {
			found = true
		}
	}
	if found {
		resp, err := http.Get("http://" + host + ":" + port + "/stopload")
		if err != nil {
			log.Errorf("err %v", err)
			return fmt.Sprintf("Err %v", err)
		}
		log.Infof("SimLoad sfu %v", resp.StatusCode)
		return "HOST_FOUND"
	} else {
		return "HOST_PORT_NOT_FOUND"
	}

}

type loadResponse struct {
	Cpu   float64 `json:"cpu"`
	Tasks int     `json:"tasks"`
	Ip    string  `json:"ip"`
	Port  string  `json:"port"`
}

type loadStatResponse struct {
	Clients     int          `json:"clients"`
	TotalRecvBW int          `json:"totalRecvBW"`
	TotalSendBW int          `json:"totalSendBW"`
	Engine      int          `json:"engine"`
	Hostload    loadResponse `json:"hostload"`
}

type statResponse struct {
	Ip    string
	Port  string
	Error string
	Stats loadStatResponse
}

func (e *etcdCoordinator) statsLoadAll() []statResponse {

	var stats []statResponse
	for _, h := range e.actionhosts {

		hstats := e.statsLoad(h.Ip, h.Port)
		stats = append(stats, hstats)
	}
	log.Infof("load stats %v", stats)
	return stats
}

func (e *etcdCoordinator) statsLoad(ip string, port string) statResponse {
	hstats := statResponse{
		Ip:   ip,
		Port: port,
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get("http://" + ip + ":" + port + "/load/stats")
	if err != nil {
		hstats.Error = fmt.Sprintf("err %v", err)
	} else {
		body, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			log.Errorf("%v", err)
			hstats.Error = fmt.Sprintf("err %v", err)
		} else {
			var response loadStatResponse
			err = json.Unmarshal(body, &response)
			if err != nil {
				log.Errorf("error parsing host response", err)
			}
			hstats.Stats = response
		}
	}
	return hstats
}

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

	apiurl := "http://" + host.Ip + ":" + host.Port + "/syncsfu/" + session + "/" + session2 + "/" + host.String() + "/" + nhost.String()
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
