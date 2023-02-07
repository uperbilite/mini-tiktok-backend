package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"mini-tiktok-backend/cmd/favorite/dal"
	favorite "mini-tiktok-backend/kitex_gen/favorite/favoriteservice"
	"net"
)

func Init() {
	dal.Init()
}

func main() {
	Init()

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	svr := favorite.NewServer(new(FavoriteServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
