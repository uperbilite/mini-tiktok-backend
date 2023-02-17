package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"mini-tiktok-backend/cmd/relation/dal"
	"mini-tiktok-backend/cmd/relation/rpc"
	"mini-tiktok-backend/kitex_gen/relation/relationservice"
	"mini-tiktok-backend/pkg/consts"
	"mini-tiktok-backend/pkg/mw"
	"net"
)

func Init() {
	dal.Init()
	rpc.Init()
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, consts.RelationServiceAddr)
	if err != nil {
		panic(err)
	}

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.RelationServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)

	Init()

	svr := relationservice.NewServer(new(RelationServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithMuxTransport(),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.RelationServiceName}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
