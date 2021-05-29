package mirrorsfu

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	connection "github.com/manishiitg/actions/connection"
	log "github.com/pion/ion-log"
	sdk "github.com/pion/ion-sdk-go"
	"github.com/pion/rtcp"
	"github.com/pion/webrtc/v3"
)

var lock sync.Mutex

type TrackMap struct {
	id    string
	track *webrtc.TrackLocalStaticRTP
}

var tracks = make(map[string][]TrackMap)

func InitWithAddress(session, session2, addr, addr2 string, cancel chan struct{}) {
	// add stun servers
	log.Warnf("InitWithAddress")
	addr = strings.Replace(addr, "700", "5005", -1)   //TODO Find a better way for this
	addr2 = strings.Replace(addr2, "700", "5005", -1) //TODO Find a better way for this
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
	// create a new client from engine
	uniq := rand.Intn(1000000)
	cid1 := fmt.Sprintf("client-mirror-1-%v", uniq)
	c1, err := sdk.NewClient(e, addr, cid1)
	if err != nil {
		log.Errorf("err=%v", err)
		return
	}

	cid2 := fmt.Sprintf("client-mirror-2-%v", uniq)
	c2, err := sdk.NewClient(e, addr2, cid2)
	if err != nil {
		log.Errorf("err=%v", err)
		return
	}

	done := make(chan struct{})

	// c1.OnDataChannel = func(dc *webrtc.DataChannel) {
	// 	go dctodc(dc, c2)
	// }
	// c2.OnDataChannel = func(dc *webrtc.DataChannel) {
	// 	go dctodc(dc, c1)
	// }
	c1.OnTrack = func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		go tracktotrack(track, receiver, c2, done, cid1)
	}
	// c2.OnTrack = func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
	// 	go tracktotrack(track, receiver, c1, done, cid2)
	// }

	err = c1.Join(session, nil)
	defer c1.Close()
	if err != nil {
		log.Errorf("err=%v", err)
		return
	}
	log.Warnf("c1 joined %v session %v", addr, session)

	err = c2.Join(session2, nil)
	defer c2.Close()
	if err != nil {
		log.Errorf("err=%v", err)
		return
	}
	log.Warnf("c2 joined %v session %v", addr2, session2)
	log.Warnf("mirroring now")
	ticker := time.NewTicker(10 * time.Second)

	go e.Stats(3, cancel)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			log.Warnf("mirror finished session %v addr1 %v session2 %v addr2 %v", session, addr, session2, addr2)
			return
		case <-ticker.C:
			lock.Lock()
			no1 := len(tracks[cid1])
			no2 := len(tracks[cid2])
			// log.Infof("session tracker c1:%v no1:%v c2:%v no2:%v", cid1, no1, cid2, no2)
			lock.Unlock()
			if no1 == 0 || no2 == 0 {
				// log.Warnf("no tracks found closing")
				// close(done)
			} else {
				// fmt.Println("tracks found! ", no)
			}

		}
	}
}

func Init(session, session2 string, cancel chan struct{}) {
	notify := make(chan string)                                                        //TODO this pattern doesn't seem proper use context with cancel etc
	go connection.GetHost("http://5.9.18.28:4000/", session, notify, cancel, "pubsub") //TODO hard coded host
	sfu_host := <-notify
	if strings.Index(sfu_host, "=") != -1 {
		session = strings.Split(sfu_host, "=")[1]
		sfu_host = strings.Split(sfu_host, "=")[0]
	}

	notify2 := make(chan string)
	go connection.GetHost("http://5.9.18.28:4000/", session2, notify2, cancel, "pubsub")
	sfu_host2 := <-notify2
	if strings.Index(sfu_host2, "=") != -1 {
		session2 = strings.Split(sfu_host2, "=")[1]
		sfu_host2 = strings.Split(sfu_host2, "=")[0]
	}
	InitWithAddress(session, session2, sfu_host, sfu_host2, cancel)
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
		return
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

func tracktotrack(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver, c2 *sdk.Client, done chan struct{}, cid string) {
	log.Warnf("GOT TRACK id%v mime%v kind %v stream %v", track.ID(), track.Codec().MimeType, track.Kind(), track.StreamID())
	newTrack, err := webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{MimeType: track.Codec().MimeType}, track.ID(), "mirror-"+track.StreamID())
	if err != nil {
		panic(err)
	}
	lock.Lock()
	tracks[cid] = append(tracks[cid], TrackMap{
		id:    track.ID(),
		track: newTrack,
	})
	lock.Unlock()

	t, err := c2.Publish(newTrack)
	if err != nil {
		log.Errorf("publish err=%v", err)
		return
	}
	go func() {
		rtcpBuf := make([]byte, 1500)
		for {
			if _, _, rtcpErr := t.Sender().Read(rtcpBuf); rtcpErr != nil {
				return
			}
		}
	}()
	defer c2.UnPublish(t)
	defer func() {
		log.Warnf("unpublish track here")
		lock.Lock()
		alltracks := tracks[cid]
		newtracks := []TrackMap{}
		for _, tr := range alltracks {
			if newTrack.ID() != tr.id {
				newtracks = append(newtracks, tr)
			}
		}
		tracks[cid] = newtracks
		lock.Unlock()
	}()
	if track.Kind() == webrtc.RTPCodecTypeVideo {
		ticker := time.NewTicker(time.Second * 2)
		defer ticker.Stop()
		rtcpSendErr := c2.GetPubTransport().GetPeerConnection().WriteRTCP([]rtcp.Packet{&rtcp.PictureLossIndication{MediaSSRC: uint32(track.SSRC())}})
		if rtcpSendErr != nil {
			fmt.Println(rtcpSendErr)
		}
	}

	for {
		select {
		case <-done:
			log.Warnf("stopping tracks publishing")
			return
		default:
			// Read
			rtpPacket, _, err := track.ReadRTP()
			if err != nil {
				log.Warnf("track read error")
				return
				// panic(err)
			}
			if err = newTrack.WriteRTP(rtpPacket); err != nil {
				log.Warnf("track write err", err)
				return
			}
		}
	}
	log.Warnf("unpublish track")
}
