package main

import (
	"flag"
	"fmt"
	"net"

	log "github.com/pion/ion-log"
	sdk "github.com/pion/ion-sdk-go"
	"github.com/pion/rtp"
	"github.com/pion/rtp/codecs"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"
	//"github.com/pion/rtcp"
)

func main() {
	// init log
	log.Init("info")

	// parse flag
	var session, addr string
	flag.StringVar(&addr, "addr", "localhost:50052", "Ion-sfu grpc addr")
	flag.StringVar(&session, "session", "test", "join session name")
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
			Level: "info",
		},
		WebRTC: sdk.WebRTCTransportConfig{
			Configuration: webrtcCfg,
		},
	}
	// new sdk engine
	e := sdk.NewEngine(config)

	// create a new client from engine
	client, err := sdk.NewClient(e, addr, "clientid2")
	if err != nil {
		log.Errorf("err=%v", err)
		return
	}

	// client join a session

	fmt.Println("joining session=%v", session)
	err = client.Join(session)
	if err != nil {
		log.Errorf("err=%v", err)
		panic(err)
	}
	videoTrack, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: "video/vp8"}, "video", "rtmp-video-pion")
	if err != nil {
		panic(err)
	}
	audioTrack, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: "audio/opus"}, "audio", "rtmp-video-pion")
	if err != nil {
		panic(err)
	}
	client.Publish(videoTrack)
	client.Publish(audioTrack)

	go rtpToTrack(videoTrack, &codecs.VP8Packet{}, 90000, 5104)
	go rtpToTrack(audioTrack, &codecs.OpusPacket{}, 48000, 5106)

	fmt.Println("waiting for rtp packets")
	select {}
}

// Listen for incoming packets on a port and write them to a Track
func rtpToTrack(track *webrtc.TrackLocalStaticSample, depacketizer rtp.Depacketizer, sampleRate uint32, port int) {
	// Open a UDP Listener for RTP Packets on port 5004
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: port})
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = listener.Close(); err != nil {
			panic(err)
		}
	}()

	sampleBuffer := samplebuilder.New(10, depacketizer, sampleRate)

	// Read RTP packets forever and send them to the WebRTC Client
	for {
		inboundRTPPacket := make([]byte, 1500) // UDP MTU
		packet := &rtp.Packet{}

		n, _, err := listener.ReadFrom(inboundRTPPacket)
		if err != nil {
			panic(fmt.Sprintf("error during read: %s", err))
		}

		if err = packet.Unmarshal(inboundRTPPacket[:n]); err != nil {
			panic(err)
		}

		sampleBuffer.Push(packet)
		for {
			sample := sampleBuffer.Pop()
			if sample == nil {
				break
			}
			if writeErr := track.WriteSample(*sample); writeErr != nil {
				panic(writeErr)
			}
		}
	}
}

// Read incoming RTCP packets
// Before these packets are retuned they are processed by interceptors. For things
// like NACK this needs to be called.
func processRTCP(rtpSender *webrtc.RTPSender) {
	go func() {
		rtcpBuf := make([]byte, 1500)

		for {
			if _, _, rtcpErr := rtpSender.Read(rtcpBuf); rtcpErr != nil {
				return
			}
		}
	}()
}
