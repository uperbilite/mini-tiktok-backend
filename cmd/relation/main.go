package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"log"
	"mini-tiktok-backend/cmd/relation/dal"
	"mini-tiktok-backend/cmd/relation/rpc"
	relation "mini-tiktok-backend/kitex_gen/relation/relationservice"
)

func Init() {
	dal.Init()
	rpc.Init()
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
}

func main() {
	svr := relation.NewServer(new(RelationServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
