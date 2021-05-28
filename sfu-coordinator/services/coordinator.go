package coordinator

import (
	"sync"
	"time"

	cloud "github.com/golivex/sfu-coordinator/cloud"

	log "github.com/pion/ion-log"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

type etcdCoordinator struct {
	mu  sync.Mutex
	cli *clientv3.Client

	hosts    map[string]Host        `json:"hosts"`
	sessions map[string]LiveSession `json:"sessions"`
	cloud    *cloud.Hub
	block    blockedSession
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
		block: blockedSession{
			session: make(map[string]string),
		},
	}

}

func (e *etcdCoordinator) RegisterCloudProvider(h *cloud.Hub) {
	e.cloud = h
}

func (e *etcdCoordinator) Close() {
	e.cli.Close()
}
