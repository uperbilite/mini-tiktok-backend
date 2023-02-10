package db

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mini-tiktok-backend/pkg/consts"
)

var DB *gorm.DB
var RDB *redis.Client

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(consts.MySQLDefaultDSN), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic(err)
	}
	RDB = redis.NewClient(&redis.Options{
		Addr:     consts.RedisAddr,
		Password: consts.RedisPassword,
		DB:       consts.RedisDB,
	})
}
