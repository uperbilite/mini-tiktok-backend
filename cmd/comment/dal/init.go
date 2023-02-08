package dal

import "mini-tiktok-backend/cmd/comment/dal/db"

// Init init dal
func Init() {
	db.Init() // mysql init
}
