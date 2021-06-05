package rtmptotrack

import (
	"fmt"
	"strings"

	"github.com/lucsky/cuid"
	connection "github.com/manishiitg/actions/connection"
	log "github.com/pion/ion-log"
	sdk "github.com/pion/ion-sdk-go"
	"github.com/pion/webrtc/v3"
)

func Init(session string, addr string, cancel <-chan struct{}) *sdk.Engine {
	webrtcCfg := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			webrtc.ICEServer{
				URLs: []string{"stun:stun.stunprotocol.org:3478", "stun:stun.l.google.com:19302"},
			},
		},
	}

	config := sdk.Config{
		WebRTC: sdk.WebRTCTransportConfig{
			Configuration: webrtcCfg,
		},
	}
	// new sdk engine
	e := sdk.NewEngine(config)

	notify := make(chan string, 1)
	go connection.GetHost(addr, session, notify, cancel, "sub", -1)
	sfu_host := <-notify

	if strings.Index(sfu_host, "=") != -1 {
		session = strings.Split(sfu_host, "=")[1]
		sfu_host = strings.Split(sfu_host, "=")[0]
	}

	// create a new client from engine
	cid := fmt.Sprintf("%s_tracktodisk_%s", session, cuid.New())
	client, err := sdk.NewClient(e, sfu_host, cid)
	if err != nil {
		log.Errorf("err=%v", err)
	}
	go run(e, client, session, cancel)
	return e
}

func run(e *sdk.Engine, client *sdk.Client, session string, cancel <-chan struct{}) {

	videoTrack, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeVP8, ClockRate: 90000}, "video", "rtmptotrack")
	if err != nil {
		panic(err)
	}

	audioTrack, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeOpus, ClockRate: 48000, Channels: 2}, "audio", "rtmptotrack")
	if err != nil {
		panic(err)
	}

	log.Infof("joining session=%v", session)
	client.Join(session, nil)
	client.Publish(audioTrack)
	client.Publish(videoTrack)
	defer e.DelClient(client)

	startRTMPServer(videoTrack, audioTrack)
	log.Infof("starting rtmp")

	select {
	case <-cancel:
		return
	}
	log.Infof("closing run")
}
