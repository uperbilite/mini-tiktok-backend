package db

import (
	"context"
	"gorm.io/gorm"
	"mini-tiktok-backend/pkg/consts"
)

type Comment struct {
	gorm.Model
	UserId     int64  `json:"user_id"`
	VideoId    int64  `json:"video_id"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

func (c *Comment) TableName() string {
	return consts.CommentTableName
}

func CreateComment(ctx context.Context, comment *Comment) (*Comment, error) {
	if err := DB.WithContext(ctx).
		Create(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func DeleteComment(ctx context.Context, id int64) error {
	return DB.WithContext(ctx).
		Where("id = ?", id).
		Delete(&Comment{}).Error
}

func GetCommentsByVideoId(ctx context.Context, videoId int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := DB.WithContext(ctx).
		Model(&Comment{}).
		Where("video_id = ?", videoId).
		Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func QueryCommentCount(ctx context.Context, videoId int64) (int64, error) {
	var res int64
	if err := DB.WithContext(ctx).
		Model(&Comment{}).
		Where("video_id = ?", videoId).
		Count(&res).Error; err != nil {
		return 0, err
	}
	return res, nil
}
