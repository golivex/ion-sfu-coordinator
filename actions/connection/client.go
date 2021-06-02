package connection

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	log "github.com/pion/ion-log"
)

type HostResponse struct {
	Host    string
	Session string
	Status  string
	Publish bool //not using this as of now as purpose of load test will fail
}

func GetHost(addr string, new_session string, notify chan string, cancel <-chan struct{}, role string, capacity int) {

	// notify <- "0.0.0.0:50052"

	done := make(chan struct{})
	for {
		select {
		case <-done:
			// log.Warnf("get host done!")
			return
		case <-cancel:
			log.Warnf("get host cancelled, cleanup")
			return
		default:
			// log.Warnf("getting sfu from %v for capacity %v", addr, capacity)
			var resp *http.Response
			var err error
			if capacity == -1 {
				resp, err = http.Get(addr + "session/" + new_session + "?role=" + role)
			} else {
				resp, err = http.Get(addr + "session/" + new_session + "?role=" + role + "&capacity=" + strconv.Itoa(capacity))
			}
			if err != nil {
				log.Errorf("%v", err)
				time.Sleep(10 * time.Second)
			} else {
				body, err2 := ioutil.ReadAll(resp.Body)
				if err2 != nil {
					log.Errorf("%v", err)
					time.Sleep(10 * time.Second)
				} else {
					var response HostResponse
					err = json.Unmarshal(body, &response)
					if err != nil {
						log.Errorf("error parsing host response", err)
					}
					sfu_host := response.Host
					log.Warnf("response %v", response, " status %v", response.Status)
					if sfu_host == "NO_HOSTS_RETRY" {
						// fmt.Println("waiting for host to get ready!")
						time.Sleep(2 * time.Second)
					} else if sfu_host == "SERVER_LOAD" {
						// fmt.Println("server is underload need to wait before joining call!")
						time.Sleep(2 * time.Second)
					} else if len(sfu_host) == 0 {
						// fmt.Println("host not found")
						time.Sleep(2 * time.Second)
					} else {
						sfu_host = strings.Replace(sfu_host, "700", "5005", -1)
						// sfu_host = strings.Replace(sfu_host, "7003", "50053", -1)
						sfu_host = strings.Replace(sfu_host, "\"", "", -1)
						if len(response.Session) > 0 {
							fmt.Println("sfu host host", sfu_host, "for session", new_session, "got new session", response.Session)
							notify <- sfu_host + "=" + response.Session //TODO this string is a temporary solution should be a strcut
						} else {
							fmt.Println("sfu host host", sfu_host, "for session", new_session)
							notify <- sfu_host
						}
						close(done)
					}
				}
			}
		}
	}

}
