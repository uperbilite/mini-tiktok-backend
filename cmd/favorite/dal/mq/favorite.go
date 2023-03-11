package mq

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"mini-tiktok-backend/pkg/consts"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	ActionType int
	UserId     int64
	VideoId    int64
}

func (msg *Message) String() string {
	return strconv.Itoa(msg.ActionType) + " " +
		strconv.FormatInt(msg.UserId, 10) + " " +
		strconv.FormatInt(msg.VideoId, 10)
}

func ParseMsg(msg string) *Message {
	ms := strings.Split(msg, " ")
	actionType, _ := strconv.Atoi(ms[0])
	userId, _ := strconv.ParseInt(ms[1], 10, 64)
	videoId, _ := strconv.ParseInt(ms[2], 10, 64)
	return &Message{
		ActionType: actionType,
		UserId:     userId,
		VideoId:    videoId,
	}
}

func (msg *Message) Produce() {
	go func() {
		for err := range Producer.Errors() {
			log.Println(err)
			// TODO: redis rollback
		}
	}()

	message := &sarama.ProducerMessage{
		Topic: consts.FavoriteTopic,
		Value: sarama.StringEncoder(msg.String()),
	}
	Producer.Input() <- message

	// TODO: wg.Wait()
}

func Consume() {
	pIds, err := Consumer.Partitions(consts.FavoriteTopic)
	if err != nil {
		panic(err)
	}

	for _, pId := range pIds {
		// create partition consumer for every partition id
		time.Sleep(20)
		pc, err := Consumer.ConsumePartition(consts.FavoriteTopic, pId, sarama.OffsetOldest)
		if err != nil {
			panic(err)
		}

		go func(pc *sarama.PartitionConsumer) {
			defer (*pc).Close()
			// block
			for message := range (*pc).Messages() {
				m := ParseMsg(string(message.Value))
				// TODO: handle msg
				fmt.Println(m.ActionType, m.VideoId, m.UserId)
				log.Printf("Partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, "111")
			}
		}(&pc)
	}
}
