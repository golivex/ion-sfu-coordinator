package cloud

import (
	"encoding/json"
	"errors"
	"os/exec"
	"time"

	"github.com/lucsky/cuid"
	log "github.com/pion/ion-log"
)

func getZone() []string {
	return []string{"asia-south1-a", "asia-south1-b", "asia-south1-c", "asia-east1-a", "asia-east1-b", "asia-east1-c", "us-central1-a", "us-central1-b"}
}

func StartInstance(zoneidx int) (machine, error) {

	var m machine
	ex := GetInstanceList()
	if len(ex) >= MAX_CLOUD_HOSTS {
		log.Infof("cannot start more hosts limit reached %v", MAX_CLOUD_HOSTS)
		return m, errors.New("cannot start more hosts")
	}

	name := cuid.New()

	if zoneidx == -1 {
		zoneidx = 0
	}
	zones := getZone()
	zone := ""
	if zoneidx < len(zones) {
		zone = zones[zoneidx]
	} else {
		return m, errors.New("all zones completed!")
	}
	name = "sfu-" + name
	output, err := exec.Command(
		"gcloud", "beta", "compute", "instances", "create", name,
		"--zone="+zone,
		"--machine-type=n1-standard-1",
		"--tags=sfu",
		"--image-family=ubuntu-2004-lts",
		"--image-project=ubuntu-os-cloud",
		"--maintenance-policy=TERMINATE",
		"--boot-disk-type=pd-ssd",
		"--metadata-from-file", "startup-script=./cloud/startup.sh",
		"--create-disk", "size=100GB,type=pd-ssd,auto-delete=yes", "--format=json").Output() //--scopes=logging-write,compute-rw,cloud-platform

	// var stdout bytes.Buffer
	// var stderr bytes.Buffer
	// cmd.Stdout = &stdout
	// cmd.Stderr = &stderr
	// err := cmd.Run()
	// log.Infof("stdout.String(), stderr.String()", stdout.String(), stderr.String())

	if err != nil {
		log.Errorf("StartServer", err)
		return StartInstance(zoneidx + 1)
	}
	// log.Debugf("output %v", string(output))

	var machines []machine
	json.Unmarshal(output, &machines)
	return machines[0], nil
}

func GetInstanceList() []machine {
	output, err := exec.Command("gcloud", "beta", "compute", "instances", "list", "--format=json").Output()
	if err != nil {
		log.Errorf("getInstanceList", err)
	}
	// var instances []map[string]interface{}
	// json.Unmarshal(output, &instances)
	// log.Infof("output %v", instances)
	// for _, value := range instances {

	// 	for key2, value2 := range value {
	// 		if key2 != "metadata" {
	// 			fmt.Println(key2, "==", value2)
	// 		}
	// 	}
	// }
	var machines []machine
	json.Unmarshal(output, &machines)

	loc, _ := time.LoadLocation("Asia/Kolkata")

	var sfum []machine
	for _, m := range machines {
		if m.isSfu() {
			m.CreationTimestamp = m.CreationTimestamp.In(loc)
			sfum = append(sfum, m)
		}
	}

	// b, _ := json.Marshal(sfum)
	// fmt.Println(string(b))
	return sfum
}

func DeleteInstance(m machine) error {
	if len(m.getName()) == 0 {
		return errors.New("Machine name cannot be empty!")
	}
	_, err := exec.Command("gcloud", "beta", "compute", "instances", "delete", m.getName(), "--zone="+m.GetZone(), "--quiet", "--format=json").Output()

	if err != nil {
		log.Errorf("DeleteServer", err)
		return err
	}
	return nil
}
