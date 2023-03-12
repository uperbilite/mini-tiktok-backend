package db

import (
	"github.com/Shopify/sarama"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mini-tiktok-backend/pkg/consts"
)

var DB *gorm.DB
var RDB *redis.Client
var Producer sarama.AsyncProducer
var Consumer sarama.Consumer

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

	initProducer()
	initConsumer()
	go Consume()
}

func initProducer() {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	var err error
	Producer, err = sarama.NewAsyncProducer([]string{consts.KafkaAddress}, config)
	if err != nil {
		panic(err)
	}
}

func initConsumer() {
	config := sarama.NewConfig()

	var err error
	Consumer, err = sarama.NewConsumer([]string{consts.KafkaAddress}, config)
	if err != nil {
		panic(err)
	}
}
