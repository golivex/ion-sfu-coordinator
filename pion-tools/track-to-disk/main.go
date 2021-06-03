package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	log "github.com/pion/ion-log"
	sdk "github.com/pion/ion-sdk-go"
	"github.com/pion/rtcp"
	"github.com/pion/rtp/codecs"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"

	//"github.com/pion/rtcp"

	"github.com/pion/rtp"

	"github.com/at-wat/ebml-go/webm"
)

type webmSaver struct {
	audioWriter, videoWriter       webm.BlockWriteCloser
	audioBuilder, videoBuilder     *samplebuilder.SampleBuilder
	audioTimestamp, videoTimestamp time.Duration
}

func newWebmSaver() *webmSaver {
	return &webmSaver{
		audioBuilder: samplebuilder.New(10, &codecs.OpusPacket{}, 48000),
		videoBuilder: samplebuilder.New(10, &codecs.VP8Packet{}, 90000),
	}
}

func (s *webmSaver) Close() {
	fmt.Printf("Finalizing webm...\n")
	if s.audioWriter != nil {
		if err := s.audioWriter.Close(); err != nil {
			panic(err)
		}
	}
	if s.videoWriter != nil {
		if err := s.videoWriter.Close(); err != nil {
			panic(err)
		}
	}
}
func (s *webmSaver) PushOpus(rtpPacket *rtp.Packet) {
	s.audioBuilder.Push(rtpPacket)

	for {
		sample := s.audioBuilder.Pop()
		if sample == nil {
			return
		}
		if s.audioWriter != nil {
			s.audioTimestamp += sample.Duration
			if _, err := s.audioWriter.Write(true, int64(s.audioTimestamp/time.Millisecond), sample.Data); err != nil {
				panic(err)
			}
		}
	}
}
func (s *webmSaver) PushVP8(rtpPacket *rtp.Packet) {
	fmt.Printf("xx %v", rtpPacket)
	s.videoBuilder.Push(rtpPacket)
	fmt.Print("yy")

	for {
		sample := s.videoBuilder.Pop()
		// fmt.Println("sample %v", sample)
		if sample == nil {
			fmt.Println("ret")
			return
		}
		// Read VP8 header.
		fmt.Print("here 76")
		videoKeyframe := (sample.Data[0]&0x1 == 0)
		if videoKeyframe {
			// Keyframe has frame information.
			raw := uint(sample.Data[6]) | uint(sample.Data[7])<<8 | uint(sample.Data[8])<<16 | uint(sample.Data[9])<<24
			width := int(raw & 0x3FFF)
			height := int((raw >> 16) & 0x3FFF)

			if s.videoWriter == nil || s.audioWriter == nil {
				// Initialize WebM saver using received frame size.
				s.InitWriter(width, height)
			}
		}
		if s.videoWriter != nil {
			s.videoTimestamp += sample.Duration
			if _, err := s.videoWriter.Write(videoKeyframe, int64(s.audioTimestamp/time.Millisecond), sample.Data); err != nil {
				panic(err)
			}
		}
		fmt.Print("loop finished")
	}
}
func (s *webmSaver) InitWriter(width, height int) {
	w, err := os.OpenFile("test.webm", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}

	ws, err := webm.NewSimpleBlockWriter(w,
		[]webm.TrackEntry{
			{
				Name:            "Audio",
				TrackNumber:     1,
				TrackUID:        12345,
				CodecID:         "A_OPUS",
				TrackType:       2,
				DefaultDuration: 20000000,
				Audio: &webm.Audio{
					SamplingFrequency: 48000.0,
					Channels:          2,
				},
			}, {
				Name:            "Video",
				TrackNumber:     2,
				TrackUID:        67890,
				CodecID:         "V_VP8",
				TrackType:       1,
				DefaultDuration: 33333333,
				Video: &webm.Video{
					PixelWidth:  uint64(width),
					PixelHeight: uint64(height),
				},
			},
		})
	if err != nil {
		panic(err)
	}
	fmt.Printf("WebM saver has started with video width=%d, height=%d\n", width, height)
	s.audioWriter = ws[0]
	s.videoWriter = ws[1]
}

func main() {
	// init log
	log.Init("debug")

	// parse flag
	var session, addr string
	flag.StringVar(&addr, "addr", "localhost:50052", "Ion-sfu grpc addr")
	flag.StringVar(&session, "session", "test2", "join session name")
	flag.Parse()

	saver := newWebmSaver()

	// add stun servers
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
	// new sdk engine
	// create a new client from engine
	client, err := sdk.NewClient(e, addr, "clientid")
	if err != nil {
		log.Errorf("err=%v", err)
		return
	}

	// subscribe rtp from sessoin
	// comment this if you don't need save to file
	client.OnTrack = func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
		if track.Kind().String() == "audio" {
			return
		}
		go func() {

			log.Infof("GOT TRACKTRACKTRACKTRACKTRACK") //, track, receiver

			// Send a PLI on an interval so that the publisher is pushing a keyframe every rtcpPLIInterval
			go func() {
				ticker := time.NewTicker(time.Second * 2)
				for range ticker.C {

					// We need to add direct access to the peerconnection to ion-sdk-go to support PLI here
					// PLI is disabled in this example currently

					if rtcpErr := client.GetSubTransport().GetPeerConnection().WriteRTCP([]rtcp.Packet{&rtcp.PictureLossIndication{MediaSSRC: uint32(track.SSRC())}}); rtcpErr != nil {
						fmt.Println(rtcpErr)
					}
				}
			}()

			for {
				// Read RTP packets being sent to Pion
				fmt.Print(track.Kind())
				rtp, _, readErr := track.ReadRTP()
				if readErr != nil {
					log.Infof("heree %v", readErr)
					if readErr == io.EOF {
						return
					}
					panic(readErr)
				}
				switch track.Kind() {
				case webrtc.RTPCodecTypeAudio:
					// saver.PushOpus(rtp)
				case webrtc.RTPCodecTypeVideo:
					saver.PushVP8(rtp)
				}
			}
			log.Infof("log tack closed")

		}()
	}
	// client join a session

	log.Infof("joining session=%v", session)
	err = client.Join(session, nil)
	if err != nil {
		log.Errorf("err=%v", err)
	}

	select {}

	saver.Close()
}
