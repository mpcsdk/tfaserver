package nats

import (
	"context"
	tfav1 "tfaserver/api/tfa/nrpc/v1"
	"tfaserver/internal/config"
	"tfaserver/internal/service"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nats-io/nats.go"
)

type sNrpcServer struct {
	sub *nats.Subscription
	nc  *nats.Conn
	//
	cache *gcache.Cache
}

var limitSendPhoneDurationCnt = 50
var limitSendPhoneDuration = time.Hour
var limitSendMailDurationCnt = 10
var limitSendMailDuration = time.Hour

func init() {
	limitSendPhoneDurationCnt = config.Config.Cache.LimitSendPhoneCount
	limitSendPhoneDuration = time.Duration(config.Config.Cache.LimitSendPhoneDuration) * time.Second
	limitSendMailDurationCnt = config.Config.Cache.LimitSendMailCount
	limitSendMailDuration = time.Duration(config.Config.Cache.LimitSendMailDuration) * time.Second

	nc, err := nats.Connect(config.Config.Nrpc.NatsUrl, nats.Timeout(5*time.Second))
	if err != nil {
		panic(err)
	}
	// defer nc.Close()
	redisCache := gcache.NewAdapterRedis(g.Redis())
	s := &sNrpcServer{
		cache: gcache.New(),
	}
	s.cache.SetAdapter(redisCache)
	///
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
