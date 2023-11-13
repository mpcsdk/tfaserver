package nats

import (
	"context"
	tfav1 "tfaserver/api/tfa/nrpc/v1"
	"tfaserver/internal/config"
	"tfaserver/internal/service"
	"time"

	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nats-io/nats.go"
)

type sNrpcServer struct {
	sub *nats.Subscription
	nc  *nats.Conn
}

func init() {
	nc, err := nats.Connect(config.Config.Nrpc.NatsUrl, nats.Timeout(5*time.Second))
	if err != nil {
		panic(err)
	}
	// defer nc.Close()

	s := &sNrpcServer{}
	h := tfav1.NewTFAHandler(gctx.GetInitCtx(), nc, s)
	sub, err := nc.QueueSubscribe(h.Subject(), "tfa", h.Handler)
	if err != nil {
		panic(err)
	}
	// defer sub.Unsubscribe()
	s.sub = sub
	s.nc = nc

	service.RegisterNrpcServer(s)
}

func (*sNrpcServer) RpcAlive(ctx context.Context, in *empty.Empty) (*empty.Empty, error) {
	//trace
	ctx, span := gtrace.NewSpan(ctx, "RpcAlive")
	defer span.End()
	//
	return &empty.Empty{}, nil
}
