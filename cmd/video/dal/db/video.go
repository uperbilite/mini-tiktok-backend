package db

import (
	"context"
	"gorm.io/gorm"
	"mini-tiktok-backend/pkg/consts"
	"strconv"
	"strings"
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

// GetVideoKey Key format is "video:{video_id}"
func GetVideoKey(videoId int64) string {
	var res strings.Builder
	res.WriteString("video:")
	res.WriteString(strconv.FormatInt(videoId, 10))
	return res.String()
}

func GetFavoriteCount(ctx context.Context, videoId int64) (int64, error) {
	res := RDB.HGet(ctx, GetVideoKey(videoId), consts.FavoriteCount)
	if res == nil {
		return 0, nil
	}
	return res.Int64()
}

func GetCommentCount(ctx context.Context, videoId int64) (int64, error) {
	res := RDB.HGet(ctx, GetVideoKey(videoId), consts.CommentCount)
	if res == nil {
		return 0, nil
	}
	return res.Int64()
}

// MGetVideos Multiple get list of videos.
func MGetVideos(ctx context.Context, videoIds []int64) ([]*Video, error) {
	res := make([]*Video, 0)
	for _, id := range videoIds {
		var v Video
		if err := DB.WithContext(ctx).
			Where("id = ?", id).
			Find(&v).Error; err != nil {
			return nil, err
		}
		res = append(res, &v)
	}
	return res, nil
}
