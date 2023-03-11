package dal

import (
	"mini-tiktok-backend/cmd/favorite/dal/db"
	"mini-tiktok-backend/cmd/favorite/dal/mq"
)

// Init init dal
func Init() {
	db.Init()
	mq.Init()
}
