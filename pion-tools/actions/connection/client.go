package connection

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	log "github.com/pion/ion-log"
)

type HostResponse struct {
	Host string
}

func GetHost(addr string, new_session string, notify chan string, cancel chan struct{}) {

	// notify <- "0.0.0.0:50052"

	done := make(chan struct{})
	for {
		select {
		case <-done:
			log.Infof("get host done!")
			return
		case <-cancel:
			log.Infof("get host cancelled, cleanup")
			return
		default:
			resp, err := http.Get(addr + "session/" + new_session + "?action=1")
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
					if sfu_host == "NO_HOSTS_RETRY" {
						fmt.Println("waiting for host to get ready!")
						time.Sleep(2 * time.Second)
					} else if sfu_host == "SERVER_LOAD" {
						fmt.Println("server is underload need to wait before joining call!")
						time.Sleep(2 * time.Second)
					} else {
						sfu_host = strings.Replace(sfu_host, "700", "5005", -1)
						// sfu_host = strings.Replace(sfu_host, "7003", "50053", -1)
						sfu_host = strings.Replace(sfu_host, "\"", "", -1)
						fmt.Println("sfu host host", sfu_host, "for session", new_session)
						notify <- sfu_host
						close(done)
					}
				}
			}
		}
	}

}
