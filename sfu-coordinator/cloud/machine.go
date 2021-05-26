package cloud

import (
	"errors"
	"strings"
	"time"
)

type machine struct {
	Id                string                           `json:"id"`
	MachineType       string                           `json:"machineType"`
	Zone              string                           `json:"zone"`
	Name              string                           `json:"name"`
	Status            string                           `json:"status"`
	CpuPlatform       string                           `json:"cpuPlatform"`
	CreationTimestamp time.Time                        `json:"creationTimestamp"`
	Tags              map[string][]string              `json:"tags"`
	NetworkInterfaces []map[string][]map[string]string `json:"networkInterfaces"`
}

func (m *machine) isSfu() bool {
	return strings.Index(m.getName(), "sfu-") != -1
}

func (m *machine) toString() string {
	return m.Name + m.GetZone()
}

func (m *machine) getName() string {
	return m.Name
}

func (m *machine) GetZone() string {
	// https://www.googleapis.com/compute/beta/projects/steady-datum-291915/zones/asia-south1-a
	split := strings.Split(m.Zone, "/")
	return split[len(split)-1]
}

func (m *machine) IsRunning() bool {
	return m.Status == "RUNNING"
}

func (m *machine) getIP() string {
	ip := ""
	for _, inf := range m.NetworkInterfaces {
		_, ok := inf["accessConfigs"]
		if ok {
			if inf["accessConfigs"][0]["name"] == "external-nat" {
				ip = inf["accessConfigs"][0]["natIP"]
				break
			}
		}
	}
	return ip
}

func (m *machine) IP() (string, error) {
	ip := ""
	for _, inf := range m.NetworkInterfaces {
		_, ok := inf["accessConfigs"]
		if ok {
			// log.Infof("net %v", inf["accessConfigs"][0]["name"])
			if inf["accessConfigs"][0]["name"] == "external-nat" {
				ip = inf["accessConfigs"][0]["natIP"]
				break
			}
		}
	}
	if len(ip) == 0 {
		return ip, errors.New("ip not found")
	} else {
		return ip, nil
	}
}
