package main

import (
	"log"
	"mini-tiktok-backend/cmd/user/dal"
	user "mini-tiktok-backend/kitex_gen/user/userservice"
)

func Init() {
	dal.Init()
}

func main() {
	Init()

	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
