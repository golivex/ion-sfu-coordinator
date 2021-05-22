package mirrorsfu

import (
	"fmt"
	"sync"
	"time"

	connection "github.com/manishiitg/actions/connection"
	log "github.com/pion/ion-log"
	sdk "github.com/pion/ion-sdk-go"
	"github.com/pion/rtcp"
	"github.com/pion/webrtc/v3"
)

var lock sync.Mutex
var tracks = make(map[string]*webrtc.TrackLocalStaticRTP)

func InitWithAddress(session, session2, addr, addr2 string) {
	// add stun servers
	webrtcCfg := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			webrtc.ICEServer{
				URLs: []string{"stun:stun.stunprotocol.org:3478", "stun:stun.l.google.com:19302"},
			},
		},
	}

	config := sdk.Config{
		Log: log.Config{
			Level: "warn",
		},
		WebRTC: sdk.WebRTCTransportConfig{
			Configuration: webrtcCfg,
		},
	}
	// new sdk engine
	e := sdk.NewEngine(config)

	// create a new client from engine
	cid1 := "client1"
	c1, err := sdk.NewClient(e, addr, cid1)
	if err != nil {
		log.Errorf("err=%v", err)
		return
	}

	cid2 := "client1"
	c2, err := sdk.NewClient(e, addr2, cid2)
	if err != nil {
		log.Errorf("err=%v", err)
		return
	}

	done := make(chan bool)

	c1.OnDataChannel = func(dc *webrtc.DataChannel) {
		go dctodc(dc, c2)
	}
	// c2.OnDataChannel = func(dc *webrtc.DataChannel) {
	// 	go dctodc(dc, c1)
	// }
	c1.OnTrack = func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		go tracktotrack(track, receiver, c2)
	}
	c2.OnTrack = func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		go tracktotrack(track, receiver, c1)
	}

	err = c1.Join(session)
	if err != nil {
		log.Errorf("err=%v", err)
		return
	}
	fmt.Println("c1 joined")

	err = c2.Join(session2)
	if err != nil {
		log.Errorf("err=%v", err)
		return
	}
	fmt.Println("c2 joined")
	fmt.Println("mirroring now")
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-done:
			fmt.Println("mirror finished session %v addr1 %v session2 %v addr2 %v", session, addr, session2, addr2)
			return
		case <-ticker.C:
			lock.Lock()
			no := len(tracks)
			lock.Unlock()
			if no == 0 {
				fmt.Println("no tracks found closing")
				close(done)
			} else {
				// fmt.Println("tracks found! ", no)
			}

		}
	}
}

func Init(session, session2 string) {
	notify := make(chan string)                                      //TODO this pattern doesn't seem proper use context with cancel etc
	go connection.GetHost("http://5.9.18.28:4000/", session, notify) //TODO hard coded host
	sfu_host := <-notify

	notify2 := make(chan string)
	go connection.GetHost("http://5.9.18.28:4000/", session2, notify2)
	sfu_host2 := <-notify2
	InitWithAddress(session, session2, sfu_host, sfu_host2)
}

func dctodc(dc *webrtc.DataChannel, c2 *sdk.Client) {
	log.Warnf("New DataChannel %s %d\n", dc.Label())
	dcID := fmt.Sprintf("dc %v", dc.Label())
	log.Warnf("DCID %v", dcID)
	dc2, err := c2.CreateDataChannel(dc.Label())
	if err != nil {
		return
	}
	dc.OnClose(func() {
		dc2.Close()
	})
	dc.OnMessage(func(msg webrtc.DataChannelMessage) {
		log.Warnf("Message from DataChannel %v %v", dc.Label(), string(msg.Data))
		dc2.SendText(string(msg.Data))
	})
	dc2.OnMessage(func(msg webrtc.DataChannelMessage) {
		// bi-directional data channels
		log.Warnf("back msg %v", string(msg.Data))
		dc.SendText(string(msg.Data))
	})
}

func tracktotrack(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver, c2 *sdk.Client) {
	log.Warnf("GOT TRACK id%v mime%v kind %v stream %v", track.ID(), track.Codec().MimeType, track.Kind(), track.StreamID())
	newTrack, err := webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{MimeType: track.Codec().MimeType}, track.ID(), track.StreamID())
	if err != nil {
		panic(err)
	}
	lock.Lock()
	tracks[track.ID()] = newTrack
	lock.Unlock()

	t, err := c2.Publish(newTrack)
	if err != nil {
		log.Errorf("publish err=%v", err)
		return
	}
	defer c2.UnPublish(t)
	defer func() {
		lock.Lock()
		delete(tracks, track.ID())
		lock.Unlock()
	}()

	go func() {
		ticker := time.NewTicker(time.Second * 2)
		for range ticker.C {
			rtcpSendErr := c2.GetSubTransport().GetPeerConnection().WriteRTCP([]rtcp.Packet{&rtcp.PictureLossIndication{MediaSSRC: uint32(track.SSRC())}})
			if rtcpSendErr != nil {
				fmt.Println(rtcpSendErr)
			}
		}
	}()

	for {
		// Read
		rtpPacket, _, err := track.ReadRTP()
		if err != nil {
			fmt.Println("track read error")
			break
			// panic(err)
		}
		if err = newTrack.WriteRTP(rtpPacket); err != nil {
			log.Errorf("track write err", err)
			break
		}
	}
	fmt.Println("unpublish track")
}
