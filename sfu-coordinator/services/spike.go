package coordinator

import (
	"time"

	log "github.com/pion/ion-log"
)

// var spikemu sync.Mutex

type Spike struct {
	Peer   int
	Tracks int
	Cpu    float64
	Time   time.Time
}

func (e *etcdCoordinator) SpikeHost(h *Host, role string) {
	//temporay skipe load on the host because it takes a few seconds i.e upto 5 - 10 for actual sfu session to start
	// need to create a temporary load till then
	// else more peers will be assigned to the host

	// log.Infof("spike host start lock")
	// spikemu.Lock()
	cpu := h.GetCurrentLoad()
	if h.PeerCount == 0 {
		cpu = cpu + 2
	} else {
		cpu = cpu + (cpu / float64(h.PeerCount))
	}
	skipe := Spike{
		Peer:   1,
		Tracks: 2,
		Cpu:    cpu,
		Time:   time.Now(),
	}

	if role == "sub" {
		cpu = cpu + 1
		skipe = Spike{
			Peer:   1,
			Tracks: 1,
			Cpu:    cpu,
			Time:   time.Now(),
		}
	}

	h.Spike = append(h.Spike, skipe)

	for key, host := range e.hosts {
		if h.String() == host.String() {
			e.hosts[key] = *h
			log.Errorf("host spiked %v", h.String())
			break
		}
	}
	// log.Infof("spike host start unlock")
	// spikemu.Unlock()

	// time.AfterFunc(5*time.Second, func() {
	// e.ClearSpikeLoad(h)
	// })
	// log.Infof("spike host end")
}

func (e *etcdCoordinator) ClearSpikeLoad(h *Host) {
	// log.Infof("clear skipe host lock start")
	// spikemu.Lock()
	// log.Infof("clear skipe host locked")
	newSpike := []Spike{}
	for _, spike := range h.Spike {
		if time.Now().Sub(spike.Time) < (5 * time.Second) {
			newSpike = append(newSpike, spike)
		} else {
			// log.Infof("clearing up spike load spikie %v host %v", spike, h.String())
		}
	}
	log.Infof("new skipe", newSpike)
	h.Spike = newSpike
	for key, host := range e.hosts {
		if host.String() == h.String() {
			e.hosts[key] = *h
			break
		}
	}
	// log.Infof("clear skipe host unlock")
	// spikemu.Unlock()
}

func (e *etcdCoordinator) GetSpikeLoad(h Host) (int, int, float64) {
	// log.Infof("get spike host start")
	e.ClearSpikeLoad(&h)
	// log.Infof("get spike host lock start")
	// spikemu.Lock()
	peers := 0
	tracks := 0
	cpu := float64(0)
	for _, spike := range h.Spike {
		peers = peers + spike.Peer
		tracks = tracks + spike.Tracks
		cpu = cpu + spike.Cpu
	}
	// log.Infof("get spike host unlokc start")
	// spikemu.Unlock()
	return peers, tracks, cpu
}
