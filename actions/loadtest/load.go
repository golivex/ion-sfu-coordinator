package loadtest

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/lucsky/cuid"
	connection "github.com/manishiitg/actions/connection"
	client "github.com/manishiitg/actions/loadtest/client"
	log "github.com/pion/ion-log"
	sdk "github.com/pion/ion-sdk-go"
	"github.com/pion/webrtc/v3"
)

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func getFileByType(file string) string {
	var filepath string
	if file == "360p" {
		filepath = "/var/tmp/Big_Buck_Bunny_4K.webm.360p.webm"
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			err := DownloadFile(filepath, "https://upload.wikimedia.org/wikipedia/commons/transcoded/c/c0/Big_Buck_Bunny_4K.webm/Big_Buck_Bunny_4K.webm.360p.webm")
			if err != nil {
				log.Infof("error downloading file %v", err)
				filepath = "test"
			}

		}
	}
	if file == "480p" {
		filepath = "/var/tmp/Big_Buck_Bunny_4K.webm.480p.webm"
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			err := DownloadFile(filepath, "https://upload.wikimedia.org/wikipedia/commons/transcoded/c/c0/Big_Buck_Bunny_4K.webm/Big_Buck_Bunny_4K.webm.480p.webm")
			if err != nil {
				log.Infof("error downloading file %v", err)
				filepath = "test"
			}
		}
	}

	if file == "720p" {
		filepath = "/var/tmp/Big_Buck_Bunny_4K.webm.720p.webm"
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			err := DownloadFile(filepath, "https://upload.wikimedia.org/wikipedia/commons/transcoded/c/c0/Big_Buck_Bunny_4K.webm/Big_Buck_Bunny_4K.webm.720p.webm")
			if err != nil {
				log.Infof("error downloading file %v", err)
				filepath = "test"
			}
		}
	}

	if file == "h264" {
		//TODO not working as of now need to debug
		filepath = "/var/tmp/Jellyfish_360_10s_1MB.mp4"
		if _, err := os.Stat(filepath); os.IsNotExist(err) {
			err := DownloadFile(filepath, "https://test-videos.co.uk/vids/jellyfish/mp4/h264/360/Jellyfish_360_10s_1MB.mp4")
			if err != nil {
				log.Infof("error downloading file %v", err)
				filepath = "test"
			}
		}
	}

	return filepath
}

func InitLoadTestApi(serverIp string, session string, clients int, role string, cycle int, rooms int, file string, capacity int, cancel chan struct{}) *sdk.Engine {
	if clients == 0 {
		clients = 1
	}
	filepath := getFileByType(file)
	return Init(filepath, serverIp, session, clients, cycle, 60*60, role, rooms, capacity, cancel)
}

func Init(file, gaddr, session string, total, cycle, duration int, role string, create_room int, capacity int, cancel chan struct{}) *sdk.Engine {
	file = getFileByType(file)
	log.Infof("filepath %v", file)
	video := true
	audio := true
	paddr := ""
	simulcast := ""

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
	go run(e, gaddr, session, file, role, total, duration, cycle, video, audio, simulcast, create_room, capacity, cancel)
	return e
}
func run(e *sdk.Engine, addr, session, file, role string, total, duration, cycle int, video, audio bool, simulcast string, create_room int, capacity int, cancel chan struct{}) *sdk.Engine {
	log.Warnf("run session=%v file=%v role=%v total=%v duration=%v cycle=%v video=%v audio=%v simulcast=%v\n", session, file, role, total, duration, cycle, audio, video, simulcast)
	timer := time.NewTimer(time.Duration(duration) * time.Second)

	go e.Stats(3, cancel)

	for i := 0; i < total; i++ {
		go func(i int, session string) {
			new_session := session
			if create_room != -1 {
				new_session = new_session + fmt.Sprintf("%v", i%create_room)
			}
			notify := make(chan string, 1)
			go connection.GetHost(addr, new_session, notify, cancel, role, capacity)
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
					var producer *client.GSTProducer
					log.Warnf("starrting new gst producer")
					if file == "test" {
						producer = client.NewGSTProducer("video", "")
					} else {
						producer = client.NewGSTProducer("screen", file)
					}
					log.Warnf("publishing tracks")

					t, _ := c.Publish(producer.VideoTrack())
					go func() {
						rtcpBuf := make([]byte, 1500)
						for {
							if _, _, rtcpErr := t.Sender().Read(rtcpBuf); rtcpErr != nil {
								log.Errorf("videoSender rtcp error", err)
								return
							}
						}
					}()
					defer c.UnPublish(t)
					defer t.Stop()
					t2, _ := c.Publish(producer.AudioTrack())
					go func() {
						rtcpBuf := make([]byte, 1500)
						for {
							if _, _, rtcpErr := t2.Sender().Read(rtcpBuf); rtcpErr != nil {
								log.Errorf("videoSender rtcp error", err)
								return
							}
						}
					}()
					defer c.UnPublish(t2)
					defer t2.Stop()
					go producer.Start()
					defer producer.Stop()
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
}
