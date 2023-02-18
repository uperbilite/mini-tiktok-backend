package dal

import (
	"mini-tiktok-backend/cmd/relation/dal/db"
)

// Init init dal
func Init() {
	db.Init() // mysql init
}
