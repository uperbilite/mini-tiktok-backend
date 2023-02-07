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

func CreateFavorite(ctx context.Context, favorites []*Favorite) error {
	return DB.WithContext(ctx).
		Create(favorites).Error
}

func DeleteFavorite(ctx context.Context, UserId int64, VideoId int64) error {
	return DB.WithContext(ctx).
		Where("user_id = ? and video_id = ? ", UserId, VideoId).
		Delete(&Favorite{}).Error
}

func QueryFavorite(ctx context.Context, UserId int64, VideoId int64) ([]*Favorite, error) {
	res := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).
		Where("user_id = ? and video_id = ? ", UserId, VideoId).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
