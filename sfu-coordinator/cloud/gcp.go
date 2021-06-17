package cloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/lucsky/cuid"
	log "github.com/pion/ion-log"
)

const DEFAULT_MACHINE_SERIES = "n1"
const DEFAULT_MACHINE_TYPE = DEFAULT_MACHINE_SERIES + "-standard-1"

func getZone() []string {
	return []string{"asia-south1-a", "asia-south1-b", "asia-south1-c", "asia-east1-a", "asia-east1-b", "asia-east1-c", "us-central1-a", "us-central1-b"}
}

// n1-standard-1 $0.04749975 almost 200 subscribers, but if 1 publishers than 150 and if 2 publishers than 100-150 subscribers
// n1-standard-2 $0.0949995
// n1-standard-4 $0.189999
// n1-standard-8 $0.379998
// n1-standard-16 $0.759996

// n1-highcpu-2	$0.0708486
// n1-highcpu-4 $0.1416972
// n1-highcpu-8 $0.2833944
// n1-highcpu-16 $0.5667888

// n2-standard-2 $0.097118
// n2-standard-4 $0.194236
// n2-standard-8 $0.388472

// n2-highcpu-2	$0.071696
// n2-highcpu-4	$0.143392
// n2-highcpu-8	$0.286784
// n2-highcpu-16 $0.573568

//n2d-standard-2 $0.084492
// n2d-standard-4 $0.168984
// n2d-highcpu-2 $0.062376
// n2d-highcpu-4 $0.124752
// machineTypes := []string{"n1-standard-1", "n1-highcpu-2", "n1-standard-4"}

func GetInstanceCapablity(mtype string) int {
	//return rough idea of instance capablity
	s := strings.Split(mtype, "-")
	cputype := s[0]
	vcpu := s[2]
	i, _ := strconv.Atoi(vcpu)
	if cputype == "n1" {
		return i * 20 //approve 20 peers per core
	}
	if cputype == "n2" {
		return i * 30 //approx 30 peers per core
	}
	return -1
}

func StartInstance(capacity int, zoneidx int, isaction bool, prefix string) (machine, error) {
	var m machine
	ex := GetInstanceList()

	sfucount := 0
	actioncount := 0
	for _, m := range ex {
		if m.isSfu() {
			sfucount = sfucount + 1
		}
		if m.isAction() {
			actioncount = actioncount + 1
		}
	}
	if sfucount >= MAX_CLOUD_HOSTS && !isaction {
		log.Infof("cannot start more sfu hosts limit reached %v sfu count %v", MAX_CLOUD_HOSTS, sfucount)
		return m, errors.New("cannot start more sfu hosts")
	}
	if actioncount >= MAX_CLOUD_ACTION_MACHINES && isaction {
		log.Infof("cannot start more action hosts limit reached %v sfu count %v", MAX_CLOUD_ACTION_MACHINES, actioncount)
		return m, errors.New("cannot start more action hosts")
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
		return m, errors.New("all zones completed")
	}
	machine_type := DEFAULT_MACHINE_SERIES + "-standard-1"
	if capacity == -1 {
		machine_type = DEFAULT_MACHINE_TYPE
	} else {
		if capacity >= 20 && capacity < 40 {
			machine_type = DEFAULT_MACHINE_SERIES + "-highcpu-2"
		} else if capacity >= 40 && capacity < 80 {
			machine_type = DEFAULT_MACHINE_SERIES + "-highcpu-4"
		} else if capacity >= 80 && capacity < 200 {
			machine_type = DEFAULT_MACHINE_SERIES + "-highcpu-8"
		} else if capacity >= 200 && capacity < 300 {
			machine_type = DEFAULT_MACHINE_SERIES + "-highcpu-16"
		} else if capacity >= 300 && capacity < 500 {
			machine_type = DEFAULT_MACHINE_SERIES + "-highcpu-32"
		} else {
			machine_type = DEFAULT_MACHINE_TYPE
		}
	}
	log.Infof("starting server with capacity %v for machine_type %v is sfu %v", capacity, machine_type, !isaction)
	if isaction {
		name = "action-" + name
	} else {
		name = "sfu-" + name
	}
	startupscript := "startup-script=./cloud/scripts/sfu-image-startup.sh"
	imagename := "sfu-minimal-image"
	tag := "sfu"
	if isaction {
		startupscript = "startup-script=./cloud/scripts/actions-image-startup.sh"
		tag = "action"
		imagename = "actions-minimal-image"
	}
	// output, err := exec.Command(
	cmd := exec.Command(
		"gcloud", "compute", "instances", "create", name,
		"--zone="+zone,
		"--machine-type="+machine_type,
		"--tags="+tag,
		"--image="+imagename,
		// "--image-family=ubuntu-minimal-2010",
		// "--image-project=ubuntu-os-cloud",
		"--maintenance-policy=TERMINATE",
		"--boot-disk-type=pd-ssd",
		"--metadata-from-file", startupscript,
		// "--metadata-from-file", "startup-script=./cloud/scripts/sfu-base-startup.sh",
		"--create-disk", "size=100GB,type=pd-ssd,auto-delete=yes", "--format=json")
	// .Output() //--scopes=logging-write,compute-rw,cloud-platform

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	output := stdout

	if err != nil {
		log.Infof("stdout.String(), stderr.String()", stdout.String(), stderr.String())
		log.Errorf("StartServer %v %v", err, output)
		return StartInstance(capacity, zoneidx+1, isaction, prefix)
	}
	// log.Debugf("output %v", output.String())

	var machines []machine
	json.Unmarshal(output.Bytes(), &machines)
	return machines[0], nil
}

func GetInstanceList() []machine {

	// cmd := exec.Command("gcloud", "compute", "instances", "list", "--format=json")

	// var stdout bytes.Buffer
	// var stderr bytes.Buffer
	// cmd.Stdout = &stdout
	// cmd.Stderr = &stderr
	// err := cmd.Run()
	// if err != nil {
	// 	log.Errorf("getInstanceList %v", err)
	// }
	// log.Infof("stdout.String(), stderr.String()", stdout.String(), stderr.String())

	output, err := exec.Command("gcloud", "compute", "instances", "list", "--format=json").Output()
	if err != nil {
		log.Errorf("getInstanceList %v", err)
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
		if (m.isSfu() || m.isAction()) && m.IsRunning() {
			m.CreationTimestamp = m.CreationTimestamp.In(loc)
			// log.Infof("found instance %v", m.toString())
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
	_, err := exec.Command("gcloud", "compute", "instances", "delete", m.getName(), "--zone="+m.GetZone(), "--quiet", "--format=json").Output()

	if err != nil {
		log.Errorf("DeleteServer %v", err)
		return err
	}
	return nil
}
