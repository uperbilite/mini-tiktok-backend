package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mini-tiktok-backend/pkg/consts"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(consts.MySQLDefaultDSN), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}
}
