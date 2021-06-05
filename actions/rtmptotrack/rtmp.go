package rtmptotrack

import (
	"bytes"
	"io"
	"net"
	"time"

	log "github.com/pion/ion-log"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/pkg/errors"
	flvtag "github.com/yutopp/go-flv/tag"
	"github.com/yutopp/go-rtmp"
	rtmpmsg "github.com/yutopp/go-rtmp/message"
)

func startRTMPServer(videoTrack, audioTrack *webrtc.TrackLocalStaticSample) {
	log.Infof("Starting RTMP Server")

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1935")
	if err != nil {
		log.Errorf("Failed: %+v", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Errorf("Failed: %+v", err)
	}

	srv := rtmp.NewServer(&rtmp.ServerConfig{
		OnConnect: func(conn net.Conn) (io.ReadWriteCloser, *rtmp.ConnConfig) {
			return conn, &rtmp.ConnConfig{
				Handler: &Handler{
					videoTrack: videoTrack,
					audioTrack: audioTrack,
				},

				ControlState: rtmp.StreamControlStateConfig{
					DefaultBandwidthWindowSize: 6 * 1024 * 1024 / 8,
				},
			}
		},
	})
	if err := srv.Serve(listener); err != nil {
		log.Errorf("Failed: %+v", err)
	}
}

type Handler struct {
	rtmp.DefaultHandler
	videoTrack, audioTrack *webrtc.TrackLocalStaticSample
}

func (h *Handler) OnServe(conn *rtmp.Conn) {
}

func (h *Handler) OnConnect(timestamp uint32, cmd *rtmpmsg.NetConnectionConnect) error {
	log.Infof("OnConnect: %#v", cmd)
	return nil
}

func (h *Handler) OnCreateStream(timestamp uint32, cmd *rtmpmsg.NetConnectionCreateStream) error {
	log.Infof("OnCreateStream: %#v", cmd)
	return nil
}

func (h *Handler) OnPublish(timestamp uint32, cmd *rtmpmsg.NetStreamPublish) error {
	log.Infof("OnPublish: %#v", cmd)

	if cmd.PublishingName == "" {
		return errors.New("PublishingName is empty")
	}
	return nil
}

func (h *Handler) OnAudio(timestamp uint32, payload io.Reader) error {
	log.Infof("OnAUdio")
	var audio flvtag.AudioData
	if err := flvtag.DecodeAudioData(payload, &audio); err != nil {
		return err
	}

	data := new(bytes.Buffer)
	if _, err := io.Copy(data, audio.Data); err != nil {
		return err
	}

	log.Infof("FLV Audio Data: Timestamp = %d, SoundFormat = %+v, SoundRate = %+v, SoundSize = %+v, SoundType = %+v, AACPacketType = %+v, Data length = %+v",
		timestamp,
		audio.SoundFormat,
		audio.SoundRate,
		audio.SoundSize,
		audio.SoundType,
		audio.AACPacketType,
		len(data.Bytes()),
	)

	return h.audioTrack.WriteSample(media.Sample{
		Data:     data.Bytes(),
		Duration: time.Duration((20/48000)*1000) * time.Millisecond,
	})
}

const headerLengthField = 4

var previousTime uint32

// annexbNALUStartCode := func() []byte { return []byte{0x00, 0x00, 0x00, 0x01} }

func (h *Handler) OnVideo(timestamp uint32, payload io.Reader) error {
	var video flvtag.VideoData
	if err := flvtag.DecodeVideoData(payload, &video); err != nil {
		return err
	}

	data := new(bytes.Buffer)
	if _, err := io.Copy(data, video.Data); err != nil {
		return err
	}

	outBuf := []byte{}
	videoBuffer := data.Bytes()
	// for offset := 0; offset < len(videoBuffer); {
	// 	bufferLength := int(binary.BigEndian.Uint32(videoBuffer[offset : offset+headerLengthField]))
	// 	if offset+bufferLength >= len(videoBuffer) {
	// 		break
	// 	}

	// 	offset += headerLengthField
	// 	outBuf = append(outBuf, []byte{0x00, 0x00, 0x00, 0x01}...)
	// 	outBuf = append(outBuf, videoBuffer[offset:offset+bufferLength]...)

	// 	offset += int(bufferLength)
	// }

	if video.FrameType == 1 {
		outBuf = append(outBuf, []byte{0x00, 0x00, 0x00, 0x01}...)
		outBuf = append(outBuf, videoBuffer[headerLengthField:]...)
	} else {
		outBuf = videoBuffer
	}

	duration := timestamp - previousTime

	log.Infof("FLV Video Data: Timestamp = %d, FrameType = %+v, CodecID = %+v, AVCPacketType = %+v, CT = %+v, Data length = %+v , Outbuf length = %+v",
		timestamp,
		video.FrameType,
		video.CodecID,
		video.AVCPacketType,
		video.CompositionTime,
		len(data.Bytes()),
		len(outBuf),
	)

	err := h.videoTrack.WriteSample(media.Sample{
		Data:     outBuf,
		Duration: time.Duration((duration/90000)*1000) * time.Millisecond,
		// Duration: time.Millisecond * 20,
	})
	if err != nil {
		panic(err)
	}
	return err
}

func (h *Handler) OnClose() {
	log.Infof("OnClose")
}
