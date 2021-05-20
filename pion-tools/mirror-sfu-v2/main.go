package main

import (
	"flag"
	"fmt"
	"sync"

	log "github.com/pion/ion-log"
	sdk "github.com/pion/ion-sdk-go"
	"github.com/pion/webrtc/v3"
)

var dcLock sync.RWMutex

type dcMap struct {
	client *sdk.Client
	dc     *webrtc.DataChannel
}

var dataMap = make(map[string]dcMap)

func main() {
	// init log
	log.Init("info")

	// parse flag
	var session, session2, addr, addr2 string
	flag.StringVar(&addr, "addr", "5.9.18.28:50052", "Ion-sfu grpc addr")
	flag.StringVar(&session, "session", "test", "join session name")

	flag.StringVar(&addr2, "addr2", "5.9.18.28:50052", "Ion-sfu grpc addr")
	flag.StringVar(&session2, "session2", "test2", "join session name")
	flag.Parse()

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
	c1, err := sdk.NewClient(e, addr, "clientid1")
	if err != nil {
		log.Errorf("err=%v", err)
		return
	}

	c2, err := sdk.NewClient(e, addr2, "clientid2")
	if err != nil {
		log.Errorf("err=%v", err)
		return
	}

	c1.OnDataChannel = func(dc *webrtc.DataChannel) {
		log.Warnf("New DataChannel %s %d\n", dc.Label())
		dcID := fmt.Sprintf("dc %v", dc.Label())
		log.Warnf("DCID %v", dcID)
		client, err := sdk.NewClient(e, addr, dcID)
		if err != nil {
			log.Errorf("err=%v", err)
			return
		}
		dcLock.Lock()
		client.Join(session2)
		dcc, err := client.CreateDataChannel(dc.Label())
		dcc.OnMessage(func(msg webrtc.DataChannelMessage) {
			// bi-directional data channels
			log.Warnf("back msg %v", string(msg.Data))
			dc.SendText(string(msg.Data))
		})
		if err != nil {
			panic(err)
		}
		dataMap[dcID] = dcMap{
			client: client,
			dc:     dcc,
		}
		dcLock.Unlock()
		dc.OnClose(func() {
			dcLock.Lock()
			defer dcLock.Unlock()
			dcID := fmt.Sprintf("closing data channel dc %v", dc.Label())
			dataMap[dcID].dc.Close()
			dataMap[dcID].client.Close()
			delete(dataMap, dcID)
		})
		dc.OnMessage(func(msg webrtc.DataChannelMessage) {
			log.Warnf("Message from DataChannel %v %v", dc.Label(), string(msg.Data))
			dcLock.Lock()
			defer dcLock.Unlock()
			dcID := fmt.Sprintf("dc %v", dc.Label())
			dataMap[dcID].dc.SendText(string(msg.Data))
		})
	}
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

	err = c2.Join(session2)
	if err != nil {
		log.Errorf("err=%v", err)
		return
	}

	select {}
}

func tracktotrack(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver, c2 *sdk.Client) {
	log.Warnf("GOT TRACK id%v mime%v kind %v stream %v", track.ID(), track.Codec().MimeType, track.Kind(), track.StreamID())
	newTrack, err := webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{MimeType: track.Codec().MimeType}, track.ID(), track.StreamID())
	if err != nil {
		panic(err)
	}

	t, err := c2.Publish(newTrack)
	if err != nil {
		log.Errorf("publish err=%v", err)
		return
	}
	defer c2.UnPublish(t)

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
