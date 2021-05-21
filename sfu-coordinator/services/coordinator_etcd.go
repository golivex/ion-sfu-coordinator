package cluster

import (
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
)

type etcdCoordinator struct {
	mu     sync.Mutex
	client *clientv3.Client
}

func newCoordinatorEtcd(host string) (*etcdCoordinator, error) {
	log.Info("creating etcd client")
	cli, err := clientv3.New(clientv3.Config{
		DialTimeout: time.Second * 3,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
		Endpoints:   []string{host}, //TODO move to config
	})
	if err != nil {
		return nil, err
	}

	log.Info("created etcdCoordinator")
	return &etcdCoordinator{
		client: cli,
	}, nil
}
