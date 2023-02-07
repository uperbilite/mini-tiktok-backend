package main

import (
	"log"
	"mini-tiktok-backend/cmd/video/dal"
	"mini-tiktok-backend/cmd/video/rpc"
	video "mini-tiktok-backend/kitex_gen/video/videoservice"
)

func Init() {
	dal.Init()
	rpc.Init()
}

func main() {
	Init()

	svr := video.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
