package coordinator

import (
	"sync"
	"time"

	log "github.com/pion/ion-log"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

type Host struct {
	Ip          string `json:"ip"`
	Port        string `json:"port"`
	PeerCount   int    `json:"peer"`
	AudioTracks int    `json:"audio"`
	VideoTracks int    `json:"video"`
	Loads       []Load
}

func (h *Host) String() string {
	return h.Ip + ":" + h.Port
}
func (h *Host) Empty() bool {
	return h.Ip == ""
}

func (h *Host) GetCurrentLoad() float64 {
	loads := h.Loads
	if len(loads) == 0 {
		return float64(100)
	}
	lastload := loads[len(loads)-1]
	return lastload.Cpu
}

type Load struct {
	Cpu float64 `json:"cpu"`
	Mem float64 `json:"mem"`
}

type Track struct {
	Id   string
	Kind string
}

type LiveSession struct {
	Name        string
	Host        string
	Port        string
	PeerCount   int
	AudioTracks int
	VideoTracks int
	Peers       []Peer
}

type Peer struct {
	Id     string
	Tracks []Track
}

type etcdCoordinator struct {
	mu  sync.Mutex
	cli *clientv3.Client

	hosts    map[string]Host        `json:"hosts"`
	sessions map[string]LiveSession `json:"sessions"`
}

func NewCoordinatorEtcd(host string) *etcdCoordinator {
	log.Infof("creating etcd client")
	cli, err := clientv3.New(clientv3.Config{
		DialTimeout: time.Second * 3,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
		Endpoints:   []string{host}, //TODO move to config
	})
	if err != nil {
		panic(err)
	}

	log.Infof("created etcdCoordinator")
	return &etcdCoordinator{
		cli:      cli,
		hosts:    make(map[string]Host),
		sessions: make(map[string]LiveSession),
	}

}
func (e *etcdCoordinator) Close() {
	e.cli.Close()
}
