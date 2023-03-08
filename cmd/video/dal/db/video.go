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
	AuthorId      int64  `json:"author_id"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	Title         string `json:"title"`
	FavoriteCount uint   `json:"favorite_count"`
	CommentCount  uint   `json:"comment_count"`
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
	ok, err := RDB.HExists(ctx, GetVideoKey(videoId), consts.FavoriteCount).Result()
	if err != nil {
		return 0, err
	}

	if ok == true { // if favorite count exists in redis
		res := RDB.HGet(ctx, GetVideoKey(videoId), consts.FavoriteCount)
		if res == nil {
			return 0, nil
		}
		return res.Int64()
	} else { // get favorite from db
		var v Video
		if err = DB.Find(&v).
			Where("id", videoId).Error; err != nil {
			return 0, nil
		}
		return int64(v.FavoriteCount), nil
	}
}

func GetCommentCount(ctx context.Context, videoId int64) (int64, error) {
	ok, err := RDB.HExists(ctx, GetVideoKey(videoId), consts.FavoriteCount).Result()
	if err != nil {
		return 0, err
	}

	if ok == true { // if favorite count exists in redis
		res := RDB.HGet(ctx, GetVideoKey(videoId), consts.FavoriteCount)
		if res == nil {
			return 0, nil
		}
		return res.Int64()
	} else { // get favorite from db
		var v Video
		if err = DB.Find(&v).
			Where("id", videoId).Error; err != nil {
			return 0, nil
		}
		return int64(v.CommentCount), nil
	}
}

// MGetVideos Multiple get list of videos.
func MGetVideos(ctx context.Context, videoIds []int64) ([]*Video, error) {
	res := make([]*Video, 0)
	// TODO: goroutine optimized
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
