package coordinator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"time"

	log "github.com/pion/ion-log"
)

func (e *etcdCoordinator) simLoad(session string, clients int, role string, cycle int, rooms int, file string) string {
	// if i am doing 20 pubsub session load is 90% on n1 machine for load testing
	// if i am doing sub only then load is 25% for 20sub, but current there is a default publisher also so 20 subs and 1pub

	no_of_machines_start := 1
	max_client_per_host := 40 //start 2 vcpu by default

	if role == "sub" {
		max_client_per_host = 200
		no_of_machines_start = int(math.Ceil(float64(clients) / float64(max_client_per_host)))
	} else {
		no_of_machines_start = int(math.Ceil(float64(clients) / float64(max_client_per_host)))
	}
	capacity := clients //start server with capacity for all nodes based on role balance will automatically start the proper instance
	if no_of_machines_start >= e.cloud.GetMaxActionMachines() {
		return fmt.Sprintf("MORE THAN %v NOT SUPPORTED AS OF NOW", e.cloud.GetMaxActionMachines())
	}
	clients_per_host := int(math.Ceil(float64(clients) / float64(max_client_per_host)))
	usedActions := make(map[string]int)
	for i := 0; i < no_of_machines_start; i++ {
		if clients > max_client_per_host {
			clients_per_host = max_client_per_host
			clients = clients - max_client_per_host
		} else {
			clients_per_host = clients
		}
		actionhost := e.getReadyActionHost()
		if actionhost != nil {
			_, ok := usedActions[actionhost.Ip]
			if ok {
				actionhost = nil
			}
		}
		if actionhost == nil {
			usedActions["CLOUD_START"+strconv.Itoa(i)] = clients_per_host
			go func() {
				notifyip := e.startActionHost(20, "loadtest") //start 2vcpu machine
				log.Infof("waiting for action machine ip")
				ip := <-notifyip
				log.Infof("got action machine ip %v", ip)
				actionhost := e.getActionHostByIp(ip)
				if actionhost == nil {
					panic("host cannot be nil!")
				}
				e.simLoadForHost(session, actionhost.Ip, actionhost.Port, clients_per_host, role, cycle, rooms, file, 5, capacity)
			}()
		} else {
			log.Infof("action host found %v", actionhost.String())
			usedActions[actionhost.Ip] = clients_per_host
			e.simLoadForHost(session, actionhost.Ip, actionhost.Port, clients_per_host, role, cycle, rooms, file, 1, capacity)
		}
	}
	b, _ := json.Marshal(usedActions)
	return string(b)

}

func (e *etcdCoordinator) simLoadForHost(session string, host string, port string, clients int, role string, cycle int, rooms int, file string, retry int, capacity int) string {

	apiurl := "http://" + host + ":" + port + "/loadtest/" + session + "?clients=" + strconv.Itoa(clients) + "&role=" + role + "&cycle=" + strconv.Itoa(cycle) + "&rooms=" + strconv.Itoa(rooms) + "&file=" + file + "&capacity=" + strconv.Itoa(capacity)
	log.Infof("load api called %v retry %v", apiurl, retry)
	resp, err := http.Get(apiurl)
	if err != nil {
		log.Errorf("%v", err)
		if retry > 1 {
			time.Sleep(5) //it takes time for host to get ready
			return e.simLoadForHost(session, host, port, clients, role, cycle, rooms, file, retry-1, capacity)

		}
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
		resp, err := http.Get("http://" + host + ":" + port + "/loadtest/stop")
		if err != nil {
			log.Errorf("err %v", err)
			return fmt.Sprintf("Err %v", err)
		}
		log.Infof("SimLoad %v", resp.StatusCode)
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
	resp, err := client.Get("http://" + ip + ":" + port + "/loadtest/stats")
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
