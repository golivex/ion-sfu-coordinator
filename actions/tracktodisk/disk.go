package tracktodisk

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/lucsky/cuid"
	connection "github.com/manishiitg/actions/connection"
	"github.com/manishiitg/actions/loadtest/client/gst"
	log "github.com/pion/ion-log"
	sdk "github.com/pion/ion-sdk-go"
	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/rtp/codecs"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media/samplebuilder"

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

func InitApi(serverip string, session string, vtype string, cancel <-chan struct{}) *sdk.Engine {
	return Init(session, serverip, vtype, cancel)
}

func Init(session string, addr string, vtype string, cancel <-chan struct{}) *sdk.Engine {
	// init log

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
	go run(e, client, session, vtype, cancel)
	return e
}

func run(e *sdk.Engine, client *sdk.Client, session string, vtype string, cancel <-chan struct{}) {
	if vtype == "gstreamer" {
		compositeSavePath := "test.mp4"

		encodePipeline := fmt.Sprintf(`
				tee name=aenctee 
				tee name=venctee
				vtee. ! queue ! vtenc_h264 ! video/x-h264,chroma-site=mpeg2 ! venctee.
				atee. ! queue ! faac ! aenctee.
		`)

		encodePipeline += fmt.Sprintf(`
				qtmux name=savemux ! queue ! filesink location=%s async=false sync=false
				venctee. ! queue ! savemux.
				aenctee. ! queue ! savemux. 
			`, compositeSavePath)
		log.Infof("saving encoded stream", "path", compositeSavePath)

		log.Infof("encoding composited stream")

		compositor := gst.NewCompositorPipeline(encodePipeline)
		compositor.Play()
		client.OnTrack = func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
			compositor.AddInputTrack(track, client.GetSubTransport().GetPeerConnection())
		}
		defer compositor.Stop()
	} else {
		saver := newWebmSaver()
		client.OnTrack = func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
			go trackToDisk(track, receiver, saver, client, cancel)
		}
		defer saver.Close()
	}

	// client join a session

	log.Infof("joining session=%v", session)
	err := client.Join(session, nil)
	defer e.DelClient(client)
	if err != nil {
		log.Errorf("err=%v", err)
	}

	select {
	case <-cancel:
		return
	}
}

func trackToDisk(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver, saver *webmSaver, client *sdk.Client, cancel <-chan struct{}) {
	log.Infof("GOT Track %v", track.Kind())

	go func() {
		ticker := time.NewTicker(time.Second * 2)
		for {
			select {
			case <-ticker.C:
				if rtcpErr := client.GetSubTransport().GetPeerConnection().WriteRTCP([]rtcp.Packet{&rtcp.PictureLossIndication{MediaSSRC: uint32(track.SSRC())}}); rtcpErr != nil {
					fmt.Println(rtcpErr)
				}
			case <-cancel:
				return
			}
		}
	}()

	for {
		// Read RTP packets being sent to Pion
		rtp, _, readErr := track.ReadRTP()
		if readErr != nil {
			if readErr == io.EOF {
				return
			}
			log.Infof("err %v", readErr)
			break
		}
		switch track.Kind() {
		case webrtc.RTPCodecTypeAudio:
			saver.PushOpus(rtp)
		case webrtc.RTPCodecTypeVideo:
			fmt.Println("video")
			saver.PushVP8(rtp)
		}
	}

}

func (s *webmSaver) Close() {
	fmt.Printf("Finalizing webm...\n")
	if s.audioWriter != nil {
		if err := s.audioWriter.Close(); err != nil {
			log.Errorf("err closing audio writer %v", err)
		}
	}
	if s.videoWriter != nil {
		if err := s.videoWriter.Close(); err != nil {
			log.Errorf("err closing video writer %v", err)
		}
	}
	fmt.Printf("Finalized webm...\n")
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
				log.Infof("error pushing opus %v", err)
			}
		}
	}
}
func (s *webmSaver) PushVP8(rtpPacket *rtp.Packet) {
	s.videoBuilder.Push(rtpPacket)

	for {
		sample := s.videoBuilder.Pop()
		if sample == nil {
			fmt.Println("nil sample")
			return
		}
		// Read VP8 header.
		videoKeyframe := (sample.Data[0]&0x1 == 0)
		fmt.Println("videoKeyframe %v", videoKeyframe)
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
	}
}
func (s *webmSaver) InitWriter(width, height int) {
	fmt.Println("init writer")
	log.Infof("webm init writer")
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
	log.Infof("WebM saver has started with video width=%d, height=%d\n", width, height)
	s.audioWriter = ws[0]
	s.videoWriter = ws[1]
}
