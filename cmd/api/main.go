// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
	"mini-tiktok-backend/cmd/api/biz/mw"
	"mini-tiktok-backend/cmd/api/biz/rpc"
	"mini-tiktok-backend/pkg/consts"
)

func Init() {
	rpc.Init()
	mw.InitJWT()
	hlog.SetLogger(hertzlogrus.NewLogger())
	hlog.SetLevel(hlog.LevelInfo)
}

func main() {
	Init()

	tracer, cfg := hertztracing.NewServerTracer()

	// 设置127.0.0.1:8080用于本地运行，设置0.0.0.0:8080用于服务器运行
	h := server.Default(server.WithHostPorts(consts.ApiServiceAddr),
		server.WithHandleMethodNotAllowed(true),
		tracer,
	)

	h.Use(hertztracing.ServerMiddleware(cfg))
	pprof.Register(h)

	register(h)
	h.Spin()
}
