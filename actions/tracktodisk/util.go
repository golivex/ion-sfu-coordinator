package tracktodisk

import (
	"fmt"
	"path"
	"time"

	avp "github.com/pion/ion-avp/pkg"
	"github.com/pion/ion-avp/pkg/elements"
	"github.com/pion/webrtc/v3"

	log "github.com/pion/ion-log"
	sdk "github.com/pion/ion-sdk-go"
	"github.com/pion/rtcp"
)

func createWebmSaver(sid, pid string) avp.Element {
	filewriter := elements.NewFileWriter(
		path.Join("./out", fmt.Sprintf("%s-%s.webm", sid, pid)),
		4096,
	)
	webm := elements.NewWebmSaver()
	webm.Attach(filewriter)
	return webm
}

func pliLoop(client *sdk.Client, track *webrtc.TrackRemote, cycle uint) {
	if cycle == 0 {
		cycle = 1000
	}

	ticker := time.NewTicker(time.Duration(cycle) * time.Millisecond)
	for range ticker.C {

		err := client.GetSubTransport().GetPeerConnection().WriteRTCP([]rtcp.Packet{&rtcp.PictureLossIndication{SenderSSRC: uint32(track.SSRC()), MediaSSRC: uint32(track.SSRC())}})
		if err != nil {
			log.Errorf("error writing pli %s", err)
		}
	}
}
