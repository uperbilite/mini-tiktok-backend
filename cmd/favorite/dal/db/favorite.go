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

func (f *Favorite) TableName() string {
	return consts.FavoriteTableName
}

// GetFavoriteKey Key format is "favorite:{video_id}"
func GetFavoriteKey(videoId int64) string {
	var res strings.Builder
	res.WriteString("favorite:")
	res.WriteString(strconv.FormatInt(videoId, 10))
	return res.String()
}

func GetFavoriteCount(ctx context.Context, videoId int64) (int64, error) {
	res := RDB.Get(ctx, GetFavoriteKey(videoId))
	if res == nil {
		return 0, nil
	}
	return res.Int64()
}

func IncrFavoriteCount(ctx context.Context, videoId int64) {
	RDB.Incr(ctx, GetFavoriteKey(videoId))
}

func DecrFavoriteCount(ctx context.Context, videoId int64) {
	favoriteCount, _ := GetFavoriteCount(ctx, videoId)
	if favoriteCount > 0 {
		RDB.Decr(ctx, GetFavoriteKey(videoId))
	}
}

func CreateFavorite(ctx context.Context, favorite *Favorite) error {
	IncrFavoriteCount(ctx, favorite.VideoId)
	return DB.WithContext(ctx).Create(favorite).Error
}

func DeleteFavorite(ctx context.Context, userId int64, videoId int64) error {
	DecrFavoriteCount(ctx, videoId)
	return DB.WithContext(ctx).
		Where("user_id = ? and video_id = ? ", userId, videoId).
		Delete(&Favorite{}).Error
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
