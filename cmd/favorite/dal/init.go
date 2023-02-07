package dal

import (
	"mini-tiktok-backend/cmd/favorite/dal/db"
)

// Init init dal
func Init() {
	db.Init() // mysql init
}
