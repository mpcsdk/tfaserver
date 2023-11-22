package cmd

import (
	"context"

	"tfaserver/internal/controller/tfa"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		g.Log().Error(r.Context(), err)
		r.Response.ClearBuffer()

		code := gcode.CodeInternalError
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    code.Code(),
			Message: code.Message(),
			Data:    nil,
		})
	}
}

func ResponseHandler(r *ghttp.Request) {
	ctx := r.GetCtx()

	r.Middleware.Next()
	// There's custom buffer content, it then exits current handler.
	g.Log().Info(ctx, r.RequestURI, r.GetBodyString())
	if r.Response.BufferLength() > 0 {
		return
	}
	var (
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	r.SetError(nil)
	if code == gcode.CodeNil {
		if err != nil {
			code = gcode.CodeInternalError
		} else {
			code = gcode.CodeOK
		}
	}
	g.Log().Info(ctx, r.RequestURI, res)
	r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    code.Code(),
		Message: code.Message(),
		Data:    res,
	})
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			s := g.Server()
			s.Use(MiddlewareErrorHandler)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(MiddlewareCORS)
				group.Middleware(ResponseHandler)
				group.Bind(
					tfa.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
