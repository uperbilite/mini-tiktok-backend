package mq

import (
	"github.com/Shopify/sarama"
	"gorm.io/gorm"
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
			// block
			for message := range (*pc).Messages() {
				time.Sleep(20)
				m := ParseMsg(string(message.Value))
				switch m.ActionType {
				case 1:
					var err error
					d := db.DB.Begin()

					if err = d.Create(&db.Favorite{
						UserId:  m.UserId,
						VideoId: m.VideoId,
					}).Error; err != nil {
						d.Rollback()
					}

					if err = d.Model(&db.Video{}).
						Where("id = ?", m.VideoId).
						Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).
						Error; err != nil {
						d.Rollback()
					}

					d.Commit()
				case 2:
					var err error
					d := db.DB.Begin()

					if err = d.
						Where("user_id = ? and video_id = ? ", m.UserId, m.VideoId).
						Delete(&db.Favorite{}).Error; err != nil {
						d.Rollback()
					}

					if err = d.Model(&db.Video{}).
						Where("id = ?", m.VideoId).
						Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).
						Error; err != nil {
						d.Rollback()
					}

					d.Commit()
				}
			}
		}(&pc)
	}
}
