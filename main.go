package main

import (
	"tfaserver/internal/config"
	_ "tfaserver/internal/packed"

	_ "tfaserver/internal/logic"
	_ "tfaserver/internal/service"

	// _ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"

	"tfaserver/internal/cmd"

	_ "tfaserver/internal/config"

	"github.com/mpcsdk/mpcCommon/rand"
)

func main() {
	ctx := gctx.GetInitCtx()
	workId := config.Config.Server.WorkId
	rand.InitIdGen(workId)
	//
	gtime.SetTimeZone("Asia/Shanghai")
	///
	// ///jaeger
	// cfg := gcfg.Instance()
	// name := cfg.MustGet(ctx, "server.name", "mpc-signer").String()
	name := config.Config.Server.Name
	// jaegerUrl, err := cfg.Get(ctx, "jaegerUrl")
	// if err != nil {
	// 	panic(err)
	// }
	jaegerUrl := config.Config.JaegerUrl
	tp, err := jaeger.Init(name, jaegerUrl)
	if err != nil {
		panic(err)
	}
	defer tp.Shutdown(ctx)

	///
	///
	cmd.Main.Run(gctx.GetInitCtx())
}
