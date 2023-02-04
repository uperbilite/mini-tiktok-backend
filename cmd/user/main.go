package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"mini-tiktok-backend/cmd/user/dal"
	user "mini-tiktok-backend/kitex_gen/user/userservice"
	"net"
)

func Init() {
	dal.Init()
}

func main() {
	Init()

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8081")
	svr := user.NewServer(new(UserServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
