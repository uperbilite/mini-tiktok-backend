package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"mini-tiktok-backend/cmd/publish/dal"
	publish "mini-tiktok-backend/kitex_gen/publish/publishservice"
	"net"
)

func Init() {
	dal.Init()
}

func main() {
	Init()

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8082") // TODO: Connect by url directly.
	svr := publish.NewServer(new(PublishServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
