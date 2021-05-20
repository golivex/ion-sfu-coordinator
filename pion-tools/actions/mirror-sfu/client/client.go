package client

import (
	"fmt"
	"sync"

	log "github.com/pion/ion-log"
	sdk "github.com/pion/ion-sdk-go"
	"github.com/pion/webrtc/v3"
	uuid "github.com/satori/go.uuid"
)

var dcLock sync.RWMutex

type dcMap struct {
	client *sdk.Client
	dc     *webrtc.DataChannel
}

var dataMap = make(map[string]dcMap)

func Init(session, addr, session2, addr2 string) {
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
	uuid, err := uuid.NewV4()
	if err == nil {
		cid1 = uuid.String()
	}
	c1, err := sdk.NewClient(e, addr, cid1)
	if err != nil {
		log.Errorf("err=%v", err)
		return
	}

	cid2 := "client1"
	uuid, err = uuid.NewV4()
	if err == nil {
		cid2 = uuid.String()
	}
	c2, err := sdk.NewClient(e, addr2, cid2)
	if err != nil {
		log.Errorf("err=%v", err)
		return
	}

	notify := make(chan string)

	c1.OnClose = func() {
		notify <- "closed"
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
	fmt.Println("c1 joined")

	err = c2.Join(session2)
	if err != nil {
		log.Errorf("err=%v", err)
		return
	}
	fmt.Println("c2 joined")

	<-notify
	fmt.Println("mirror finished session %v addr1 %v session2 %v addr2 %v", session, addr, session2, addr2)

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
