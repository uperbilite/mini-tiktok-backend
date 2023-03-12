package db

import (
	"context"
	"gorm.io/gorm"
	"mini-tiktok-backend/pkg/consts"
	"strconv"
	"strings"
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

func CreateFavoriteInRedis(ctx context.Context, videoId int64) error {
	// TODO: error handle
	RDB.Incr(ctx, GetVideoKey(videoId))
	return nil
}

func CreateFavoriteInMysql(userId int64, videoId int64) {
	var err error
	d := DB.Begin()

	if err = d.Create(&Favorite{
		UserId:  userId,
		VideoId: videoId,
	}).Error; err != nil {
		d.Rollback()
	}

	if err = d.Model(&Video{}).
		Where("id = ?", videoId).
		Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).
		Error; err != nil {
		d.Rollback()
	}

	// TODO: delete key in redis if failed

	d.Commit()
}

func DeleteFavoriteInRedis(ctx context.Context, videoId int64) error {
	// TODO: error handle
	RDB.HDel(ctx, GetVideoKey(videoId), consts.FavoriteCount)
	return nil
}

func DeleteFavoriteInMysql(userId int64, videoId int64) {
	var err error
	db := DB.Begin()

	if err = db.Where("user_id = ? and video_id = ? ", userId, videoId).
		Delete(&Favorite{}).Error; err != nil {
		db.Rollback()
	}

	if err = db.Model(&Video{}).
		Where("id = ?", videoId).
		Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).
		Error; err != nil {
		db.Rollback()
	}

	db.Commit()
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
