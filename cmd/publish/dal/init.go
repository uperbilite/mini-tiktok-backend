package dal

import (
	"mini-tiktok-backend/cmd/publish/dal/db"
)

// Init init dal
func Init() {
	db.Init() // mysql init
}
