package nrpcclient

import (
	riskv1 "tfaserver/api/risk/nrpc/v1"
	v1 "tfaserver/api/risk/nrpc/v1"
	// tfav1 "tfaserver/api/tfa/nrpc/v1"
	"tfaserver/internal/config"
	"tfaserver/internal/service"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nats-io/nats.go"
)

type sNrpcClient struct {
	cli *riskv1.RiskClient
	nc  *nats.Conn
}

func init() {

	// Connect to the NATS server.
	nc, err := nats.Connect(config.Config.Nrpc.NatsUrl, nats.Timeout(3*time.Second))
	if err != nil {
		panic(err)
	}
	// defer nc.Close()

	// This is our generated client.
	cli := riskv1.NewRiskClient(nc)

	// Contact the server and print out its response.
	_, err = cli.RpcAlive(&empty.Empty{})
	if err != nil {
		panic(err)
	}
	s := &sNrpcClient{
		cli: cli,
		nc:  nc,
	}
	service.RegisterNrpcClient(s)
}
func (s *sNrpcClient) Flush() {
	err := s.nc.Flush()
	if err != nil {
		panic(err)
	}
	s.cli = v1.NewRiskClient(s.nc)
}
