package mq

import (
	"github.com/Shopify/sarama"
	"mini-tiktok-backend/pkg/consts"
)

var Producer sarama.AsyncProducer
var Consumer sarama.Consumer

func Init() {
	initProducer()
	initConsumer()
	go Consume()
}

func initProducer() {
	config := sarama.NewConfig()
	//config.Producer.Return.Successes = true
	//config.Producer.Return.Errors = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	var err error
	Producer, err = sarama.NewAsyncProducer([]string{consts.KafkaAddress}, config)
	if err != nil {
		panic(err)
	}
}

func initConsumer() {
	config := sarama.NewConfig()
	//config.Consumer.Return.Errors = true

	var err error
	Consumer, err = sarama.NewConsumer([]string{consts.KafkaAddress}, config)
	if err != nil {
		panic(err)
	}
}
