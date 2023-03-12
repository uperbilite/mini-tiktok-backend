package mq

import (
	"github.com/Shopify/sarama"
	"mini-tiktok-backend/cmd/favorite/dal/db"
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
	message := &sarama.ProducerMessage{
		Topic: consts.FavoriteTopic,
		Value: sarama.StringEncoder(msg.String()),
	}
	Producer.Input() <- message
}

func Consume() {
	pIds, err := Consumer.Partitions(consts.FavoriteTopic)
	if err != nil {
		panic(err)
	}

	for _, pId := range pIds {
		// create partition consumer for every partition id
		pc, err := Consumer.ConsumePartition(consts.FavoriteTopic, pId, sarama.OffsetOldest)
		if err != nil {
			panic(err)
		}

		go func(pc *sarama.PartitionConsumer) {
			defer (*pc).Close()
			for message := range (*pc).Messages() {
				// sleep 20ms to gentle to peak of flow
				// TODO: update mysql in batch.
				time.Sleep(20)
				m := ParseMsg(string(message.Value))
				switch m.ActionType {
				case 1:
					db.CreateFavoriteInMysql(m.UserId, m.VideoId)
				case 2:
					db.DeleteFavoriteInMysql(m.UserId, m.VideoId)
				}
			}
		}(&pc)
	}
}
