package db

import (
	"context"
	"gorm.io/gorm"
	"mini-tiktok-backend/pkg/consts"
)

type Video struct {
	gorm.Model
	AuthorId int64  `json:"author_id"`
	PlayURL  string `json:"play_url"`
	CoverURL string `json:"cover_url"`
	Title    string `json:"title"`
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}

// MGetVideos Multiple get list of videos.
func MGetVideos(ctx context.Context, videoIds []int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).
		Where("id in ?", videoIds).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
