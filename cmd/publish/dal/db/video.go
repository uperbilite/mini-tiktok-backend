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

// CreateVideo Create video.
func CreateVideo(ctx context.Context, video *Video) error {
	return DB.WithContext(ctx).Create(video).Error
}

// GetVideoIdsByAuthorId Get all videos' id by same author.
func GetVideoIdsByAuthorId(ctx context.Context, id int64) ([]int64, error) {
	res := make([]int64, 0)
	if err := DB.WithContext(ctx).
		Model(&Video{}).
		Select("id").
		Where("author_id = ?", id).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetVideoIdsFeed Get video's id feed.
func GetVideoIdsFeed(ctx context.Context, latestTime int64) ([]int64, error) {
	res := make([]int64, 0)
	if err := DB.WithContext(ctx).
		Model(&Video{}).
		Select("id").
		Where("unix_timestamp(created_at) < ?", latestTime).
		Limit(30).
		Order("created_at desc").
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
