package dal

import (
	"mini-tiktok-backend/cmd/video/dal/db"
)

// Init init dal
func Init() {
	db.Init() // mysql init
}
