package db

import (
	"context"
	"gorm.io/gorm"
	"mini-tiktok-backend/pkg/consts"
)

type Video struct {
	gorm.Model
	AuthorId       int64  `json:"author_id"`
	PlayURL        string `json:"play_url"`
	CoverURL       string `json:"cover_url"`
	FavouriteCount int    `json:"favourite_count"`
	CommentCount   int    `json:"comment_count"`
	Title          string `json:"title"`
}

func (v *Video) TableName() string {
	return consts.VideoTableName
}

// CreateVideo Create video.
func CreateVideo(ctx context.Context, video *Video) error {
	return DB.WithContext(ctx).Create(video).Error
}

// MGetVideo Get all videos by same author.
func MGetVideo(ctx context.Context, id int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).
		Where("author_id = ?", id).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetVideoFeed(ctx context.Context, latestTime int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).
		Where("unix_timestamp(created_at) < ?", latestTime).
		Limit(30).
		Order("created_at desc").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
