package loadtest

import (
	"fmt"
	"strings"
	"time"

	"github.com/lucsky/cuid"
	connection "github.com/manishiitg/actions/connection"
	client "github.com/manishiitg/actions/loadtest/client"
	ilog "github.com/pion/ion-log"
	sdk "github.com/pion/ion-sdk-go"
	"github.com/pion/webrtc/v3"
	"github.com/sirupsen/logrus"
)

var (
	log = ilog.NewLoggerWithFields(ilog.WarnLevel, "", nil)
)

func InitApi(session string, clients int, cancel chan struct{}) {
	if clients == 0 {
		clients = 1
	}
	Init("test", "http://0.0.0.0:4000/", session, clients, 1000, 60*60, "pubsub", true, true, "", "", -1, cancel, true)
	//TODO
	// tried os.exec also didn't work its not killing the process at all
	// starting os.exec here because when we stop load test the session should get closed instantly buts its not getting closed
	//even after trying a lot
	// cmd := exec.Command("go", "run", "cmd/load.go", "-session", session, "-clients", strconv.Itoa(clients))
	// err := cmd.Start()
	// if err != nil {
	// 	log.Infof("error in starting process %v", err)
	// }
	// log.Infof("process id %v", cmd.Process.Pid)
	// for {
	// 	select {
	// 	case <-cancel:
	// 		log.Infof("process cancel called")
	// 		err := cmd.Process.Signal(syscall.SIGTERM)
	// 		if err != nil {
	// 			log.Infof("Unable to kill process %v", err)
	// 		}
	// 		return
	// 	}

	// }
}

func Init(file, gaddr, session string, total, cycle, duration int, role string, video bool, audio bool, simulcast string, paddr string, create_room int, cancel chan struct{}, fromApi bool) {
	log.SetLevel(logrus.WarnLevel)
	se := webrtc.SettingEngine{}
	se.SetEphemeralUDPPortRange(10000, 15000)
	webrtcCfg := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			webrtc.ICEServer{
				URLs: []string{"stun:stun.stunprotocol.org:3478", "stun:stun.l.google.com:19302"},
			},
		},
	}
	var videoMime string = "video/vp8"
	if !strings.Contains(file, ".webm") {
		videoMime = "video/h264"
	}
	config := sdk.Config{
		WebRTC: sdk.WebRTCTransportConfig{
			VideoMime:     videoMime,
			Setting:       se,
			Configuration: webrtcCfg,
		},
	}
	if gaddr == "" {
		log.Errorf("gaddr is \"\"!")
		return
	}
	e := sdk.NewEngine(config)
	if paddr != "" {
		go e.ServePProf(paddr)
	}
	run(e, gaddr, session, file, role, total, duration, cycle, video, audio, simulcast, create_room, cancel, fromApi)
}

func run(e *sdk.Engine, addr, session, file, role string, total, duration, cycle int, video, audio bool, simulcast string, create_room int, cancel chan struct{}, fromApi bool) {
	log.Warnf("run session=%v file=%v role=%v total=%v duration=%v cycle=%v video=%v audio=%v simulcast=%v\n", session, file, role, total, duration, cycle, audio, video, simulcast)
	timer := time.NewTimer(time.Duration(duration) * time.Second)

	// defer recoverClose() //TODO see if this can be fixed from ion-sdk-go or debug how its happening https://github.com/pion/ion-sdk-go/issues/34
	if !fromApi {
		go e.Stats(3)
	}

	notify := make(chan string)
	for i := 0; i < total; i++ {
		new_session := session
		if create_room != -1 {
			new_session = new_session + fmt.Sprintf("%v", i%create_room)
		}

		go connection.GetHost(addr, new_session, notify, cancel)
		sfu_host := <-notify

		if strings.Index(sfu_host, "=") != -1 {
			new_session = strings.Split(sfu_host, "=")[1]
			sfu_host = strings.Split(sfu_host, "=")[0]
		}

		switch role {
		case "pubsub":
			var producer *client.GSTProducer
			cid := fmt.Sprintf("%s_pubsub_%d_%s", new_session, i, cuid.New())
			log.Errorf("AddClient session=%v clientid=%v addr=%v", new_session, cid, addr)
			c, err := sdk.NewClient(e, sfu_host, cid)
			if err != nil {
				log.Errorf("%v", err)
				break
			}
			c.Join(new_session, nil)
			defer e.DelClient(c)
			if !strings.Contains(file, ".webm") {
				log.Warnf("starrting new gst producer")
				if file == "test" {
					producer = client.NewGSTProducer("video", "")
				} else {
					producer = client.NewGSTProducer("screen", file)
				}
				log.Warnf("publishing tracks")
				go producer.Start()
				defer producer.Stop()
				t, _ := c.Publish(producer.VideoTrack())
				defer c.UnPublish(t)
				defer t.Stop()
				t2, _ := c.Publish(producer.AudioTrack())
				defer c.UnPublish(t2)
				defer t2.Stop()
				log.Warnf("tracks published")
			} else {
				c.PublishWebm(file, video, audio)
			}
			c.Simulcast(simulcast)
		case "sub":
			cid := fmt.Sprintf("%s_sub_%d_%s", new_session, i, cuid.New())
			log.Errorf("AddClient session=%v clientid=%v addr=%v", new_session, cid, addr)
			c, err := sdk.NewClient(e, sfu_host, cid)
			if err != nil {
				log.Errorf("%v", err)
				break
			}
			config := sdk.NewJoinConfig().SetNoPublish()
			c.Join(new_session, config)
			c.Simulcast(simulcast)
			defer e.DelClient(c)
		case "pub":
			cid := fmt.Sprintf("%s_pub_%d_%s", session, i, cuid.New())
			log.Errorf("AddClient session=%v clientid=%v addr=%v", session, cid, addr)
			c, err := sdk.NewClient(e, addr, cid)
			if err != nil {
				log.Errorf("%v", err)
				break
			}
			config := sdk.NewJoinConfig().SetNoSubscribe()
			c.Join(session, config)
			c.Simulcast(simulcast)
			defer e.DelClient(c)
			if !strings.Contains(file, ".webm") {
				var producer *client.GSTProducer
				log.Warnf("starrting new gst producer")
				if file == "test" {
					producer = client.NewGSTProducer("video", "")
				} else {
					producer = client.NewGSTProducer("screen", file)
				}
				log.Warnf("publishing tracks")
				go producer.Start()
				defer producer.Stop()
				t, _ := c.Publish(producer.VideoTrack())
				defer c.UnPublish(t)
				defer t.Stop()
				t2, _ := c.Publish(producer.AudioTrack())
				defer c.UnPublish(t2)
				defer t2.Stop()
				log.Warnf("tracks published")
			} else {
				c.PublishWebm(file, video, audio)
			}
		default:
			log.Errorf("invalid role! should be pubsub/sub")
		}

		time.Sleep(time.Millisecond * time.Duration(cycle))
	}

	select {
	case <-timer.C:
		return
	case <-cancel:
		log.Warnf("cancel called on load test")
		return
	}

}
