package db

import (
	"context"
	"gorm.io/gorm"
	"mini-tiktok-backend/pkg/consts"
)

type Favorite struct {
	gorm.Model
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

func (f *Favorite) TableName() string {
	return consts.FavoriteTableName
}

func CreateFavorite(ctx context.Context, favorite *Favorite) error {
	return DB.WithContext(ctx).Create(favorite).Error
}

func DeleteFavorite(ctx context.Context, userId int64, videoId int64) error {
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

func QueryFavoriteCount(ctx context.Context, videoId int64) (int64, error) {
	var res int64
	if err := DB.WithContext(ctx).
		Model(&Favorite{}).
		Where("video_id = ?", videoId).
		Count(&res).Error; err != nil {
		return 0, err
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
