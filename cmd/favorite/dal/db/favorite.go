package db

import (
	"context"
	"github.com/Shopify/sarama"
	"gorm.io/gorm"
	"mini-tiktok-backend/pkg/consts"
	"strconv"
	"strings"
	"time"
)

type Favorite struct {
	gorm.Model
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

type Video struct {
	gorm.Model
	AuthorId      int64  `json:"author_id"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	Title         string `json:"title"`
	FavoriteCount uint   `json:"favorite_count"`
	CommentCount  uint   `json:"comment_count"`
}

func (f *Favorite) TableName() string {
	return consts.FavoriteTableName
}

// GetVideoKey Key format is "video:{video_id}"
func GetVideoKey(videoId int64) string {
	var res strings.Builder
	res.WriteString("video:")
	res.WriteString(strconv.FormatInt(videoId, 10))
	return res.String()
}

func CreateFavorite(ctx context.Context, favorite *Favorite) error {
	var err error
	if err = RDB.Incr(ctx, GetVideoKey(favorite.VideoId)).Err(); err != nil {
		// TODO: handle error
	}

	msg := &Message{
		ActionType: 1,
		UserId:     favorite.UserId,
		VideoId:    favorite.VideoId,
	}
	msg.Produce()

	return err
}

func DeleteFavorite(ctx context.Context, userId int64, videoId int64) error {
	var err error
	db := DB.Begin()

	if err = db.WithContext(ctx).
		Where("user_id = ? and video_id = ? ", userId, videoId).
		Delete(&Favorite{}).Error; err != nil {
		db.Rollback()
	}

	if err = db.WithContext(ctx).
		Model(&Video{}).
		Where("id = ?", videoId).
		Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).
		Error; err != nil {
		db.Rollback()
	}

	db.Commit()

	// delete redis key for consistent
	if err = RDB.HDel(ctx, GetVideoKey(videoId), consts.FavoriteCount).Err(); err != nil {
		db.Rollback()
	}

	// TODO: kafka listen binlog and send to redis

	return err
}

func QueryFavorite(ctx context.Context, userId int64, videoId int64) ([]*Favorite, error) {
	res := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).
		Where("user_id = ? and video_id = ? ", userId, videoId).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetVideoIdsByUserId(ctx context.Context, id int64) ([]int64, error) {
	res := make([]int64, 0)
	if err := DB.WithContext(ctx).
		Model(&Favorite{}).
		Select("video_id").
		Where("user_id = ?", id).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

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
					db := DB.Begin()

					if err = db.Create(&Favorite{
						UserId:  m.UserId,
						VideoId: m.VideoId,
					}).Error; err != nil {
						db.Rollback()
					}

					if err = db.Model(&Video{}).
						Where("id = ?", m.VideoId).
						Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).
						Error; err != nil {
						db.Rollback()
					}

					db.Commit()
				case 2:
					var err error
					db := DB.Begin()

					if err = db.
						Where("user_id = ? and video_id = ? ", m.UserId, m.VideoId).
						Delete(&Favorite{}).Error; err != nil {
						db.Rollback()
					}

					if err = db.Model(&Video{}).
						Where("id = ?", m.VideoId).
						Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).
						Error; err != nil {
						db.Rollback()
					}

					db.Commit()
				}
			}
		}(&pc)
	}
}
