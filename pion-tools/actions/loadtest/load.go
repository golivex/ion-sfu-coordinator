package loadtest

import (
	"fmt"
	"strings"
	"time"

	connection "github.com/manishiitg/actions/connection"
	client "github.com/manishiitg/actions/loadtest/client"
	log "github.com/pion/ion-log"
	sdk "github.com/pion/ion-sdk-go"
	"github.com/pion/webrtc/v3"
)

func InitApi(session string, cancel chan struct{}) {
	Init("test", "http://0.0.0.0:4000/", session, 1, 1000, 60*5, "pubsub", true, true, "", "", -1, cancel, true)
}

func Init(file, gaddr, session string, total, cycle, duration int, role string, video bool, audio bool, simulcast string, paddr string, create_room int, cancel chan struct{}, fromApi bool) {
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
		Log: log.Config{
			Level: "warn",
		},
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

func recoverClose() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func run(e *sdk.Engine, addr, session, file, role string, total, duration, cycle int, video, audio bool, simulcast string, create_room int, cancel chan struct{}, fromApi bool) {
	log.Infof("run session=%v file=%v role=%v total=%v duration=%v cycle=%v video=%v audio=%v simulcast=%v\n", session, file, role, total, duration, cycle, audio, video, simulcast)
	timer := time.NewTimer(time.Duration(duration) * time.Second)

	defer recoverClose() //TODO see if this can be fixed from ion-sdk-go or debug how its happening

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

		switch role {
		case "pubsub":
			var producer *client.GSTProducer
			cid := fmt.Sprintf("%s_pubsub_%d", new_session, i)
			fmt.Println("AddClient session=%v clientid=%v", new_session, cid)
			c, err := sdk.NewClient(e, sfu_host, cid)
			if err != nil {
				log.Errorf("%v", err)
				break
			}
			c.Join(new_session)
			if !strings.Contains(file, ".webm") {
				fmt.Println("starrting new gst producer")
				if file == "test" {
					producer = client.NewGSTProducer("video", "")
				} else {
					producer = client.NewGSTProducer("screen", file)
				}
				log.Infof("publishing tracks")
				go producer.Start()
				t, _ := c.Publish(producer.VideoTrack())
				defer c.UnPublish(t)
				t2, _ := c.Publish(producer.AudioTrack())
				defer c.UnPublish(t2)
				log.Infof("tracks published")
			} else {
				c.PublishWebm(file, video, audio)
			}
			c.Simulcast(simulcast)
			defer c.Close()
		case "sub":
			cid := fmt.Sprintf("%s_sub_%d", new_session, i)
			log.Infof("AddClient session=%v clientid=%v", new_session, cid)
			c, err := sdk.NewClient(e, sfu_host, cid)
			if err != nil {
				log.Errorf("%v", err)
				break
			}
			c.Join(new_session)
			c.Simulcast(simulcast)
			defer c.Close()
		default:
			log.Errorf("invalid role! should be pubsub/sub")
		}

		time.Sleep(time.Millisecond * time.Duration(cycle))
	}

	select {
	case <-timer.C:
		return
	case <-cancel:
		return
	}

}
