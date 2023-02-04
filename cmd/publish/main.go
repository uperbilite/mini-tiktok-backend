package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	publish "mini-tiktok-backend/kitex_gen/publish/publishservice"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8082")
	svr := publish.NewServer(new(PublishServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
