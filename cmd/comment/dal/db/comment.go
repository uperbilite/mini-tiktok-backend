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
	if err := DB.WithContext(ctx).Create(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func DeleteComment(ctx context.Context, id int64) error {
	return DB.WithContext(ctx).
		Where("id = ?", id).
		Delete(&Comment{}).Error
}
