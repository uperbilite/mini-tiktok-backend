package db

import (
	"context"
	"gorm.io/gorm"
	"mini-tiktok-backend/pkg/consts"
	"strconv"
	"strings"
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

// GetCommentKey Key format is "comment:{video_id}"
func GetCommentKey(videoId int64) string {
	var res strings.Builder
	res.WriteString("comment:")
	res.WriteString(strconv.FormatInt(videoId, 10))
	return res.String()
}

func GetCommentCount(ctx context.Context, videoId int64) (int64, error) {
	res := RDB.Get(ctx, GetCommentKey(videoId))
	if res == nil {
		return 0, nil
	}
	return res.Int64()
}

func IncrCommentCount(ctx context.Context, videoId int64) {
	RDB.Incr(ctx, GetCommentKey(videoId))
}

func DecrCommentCount(ctx context.Context, videoId int64) {
	commentCount, _ := GetCommentCount(ctx, videoId)
	if commentCount > 0 {
		RDB.Decr(ctx, GetCommentKey(videoId))
	}
}

func CreateComment(ctx context.Context, comment *Comment) (*Comment, error) {
	if err := DB.WithContext(ctx).
		Create(comment).Error; err != nil {
		return nil, err
	}
	IncrCommentCount(ctx, comment.VideoId)
	return comment, nil
}

func DeleteComment(ctx context.Context, id int64) error {
	var videoId int64
	if err := DB.WithContext(ctx).
		Model(&Comment{}).
		Select("video_id").
		Where("id = ?", id).
		Find(&videoId).Error; err != nil {
		return err
	}
	DecrCommentCount(ctx, videoId)
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
