package dal

import (
	"mini-tiktok-backend/cmd/user/dal/db"
)

// Init init dal
func Init() {
	db.Init() // mysql init
}
