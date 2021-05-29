package loadtest

import (
	"fmt"
	"strings"
	"time"

	"github.com/lucsky/cuid"
	connection "github.com/manishiitg/actions/connection"
	client "github.com/manishiitg/actions/loadtest/client"
	log "github.com/pion/ion-log"
	sdk "github.com/pion/ion-sdk-go"
	"github.com/pion/webrtc/v3"
)

func InitLoadTestApi(serverIp string, session string, clients int, role string, cycle int, rooms int, cancel chan struct{}) *sdk.Engine {
	if clients == 0 {
		clients = 1
	}
	return Init("./big-buck-bunny_trailer.webm", "http://"+serverIp+":4000/", session, clients, cycle, 60*60, role, true, true, "", "", rooms, cancel)
}

func Init(file, gaddr, session string, total, cycle, duration int, role string, video bool, audio bool, simulcast string, paddr string, create_room int, cancel chan struct{}) *sdk.Engine {
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
	e := sdk.NewEngine(config)
	if paddr != "" {
		go e.ServePProf(paddr)
	}
	go run(e, gaddr, session, file, role, total, duration, cycle, video, audio, simulcast, create_room, cancel)
	return e
}
func run(e *sdk.Engine, addr, session, file, role string, total, duration, cycle int, video, audio bool, simulcast string, create_room int, cancel chan struct{}) *sdk.Engine {
	log.Warnf("run session=%v file=%v role=%v total=%v duration=%v cycle=%v video=%v audio=%v simulcast=%v\n", session, file, role, total, duration, cycle, audio, video, simulcast)
	timer := time.NewTimer(time.Duration(duration) * time.Second)

	// defer recoverClose() //TODO see if this can be fixed from ion-sdk-go or debug how its happening https://github.com/pion/ion-sdk-go/issues/34

	go e.Stats(3, cancel)

	// var wg sync.WaitGroup
	for i := 0; i < total; i++ {
		// wg.Add(1)
		go func(i int, session string) {
			// defer wg.Done()
			new_session := session
			if create_room != -1 {
				new_session = new_session + fmt.Sprintf("%v", i%create_room)
			}
			notify := make(chan string, 1)

			go connection.GetHost(addr, new_session, notify, cancel, role)
			sfu_host := <-notify

			if strings.Index(sfu_host, "=") != -1 {
				new_session = strings.Split(sfu_host, "=")[1]
				sfu_host = strings.Split(sfu_host, "=")[0]
			}

			crole := role

			if role == "sub" && i == 0 {
				crole = "pubsub"
			}

			switch crole {
			case "pubsub":
				var producer *client.GSTProducer
				cid := fmt.Sprintf("%s_pubsub_%d_%s", new_session, i, cuid.New())
				log.Errorf("AddClient session=%v clientid=%v addr=%v", new_session, cid, sfu_host)
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
				log.Errorf("AddClient session=%v clientid=%v addr=%v", new_session, cid, sfu_host)
				c, err := sdk.NewClient(e, sfu_host, cid)
				if err != nil {
					log.Errorf("%v", err)
					break
				}
				// config := sdk.NewJoinConfig().SetNoPublish() //TODO bug raise wait for fix
				c.Join(new_session, nil)
				defer e.DelClient(c)
				c.Simulcast(simulcast)
			case "pub":
				cid := fmt.Sprintf("%s_pub_%d_%s", session, i, cuid.New())
				log.Errorf("AddClient session=%v clientid=%v addr=%v", session, cid, sfu_host)
				c, err := sdk.NewClient(e, addr, cid)
				if err != nil {
					log.Errorf("%v", err)
					break
				}
				config := sdk.NewJoinConfig().SetNoSubscribe()
				c.Join(session, config)
				defer e.DelClient(c)
				c.Simulcast(simulcast)
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

			select {
			case <-timer.C:
				return
			case <-cancel:
				log.Warnf("cancel called on load test")
				return
			}
		}(i, session)
		time.Sleep(time.Millisecond * time.Duration(cycle))
	}
	return e
	// wg.Wait()

}