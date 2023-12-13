package nrpcclient

import (
	"tfaserver/api/risk/nrpc"

	"tfaserver/internal/config"
	"tfaserver/internal/service"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nats-io/nats.go"
)

type sNrpcClient struct {
	cli *nrpc.RiskClient
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
	cli := nrpc.NewRiskClient(nc)

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
	s.cli = nrpc.NewRiskClient(s.nc)
}
