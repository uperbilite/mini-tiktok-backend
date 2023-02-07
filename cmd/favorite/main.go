package main

import (
	"log"
	"mini-tiktok-backend/cmd/favorite/dal"
	favorite "mini-tiktok-backend/kitex_gen/favorite/favoriteservice"
)

func Init() {
	dal.Init()
}

func main() {
	Init()

	svr := favorite.NewServer(new(FavoriteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
