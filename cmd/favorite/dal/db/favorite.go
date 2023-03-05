package db

import (
	"context"
	"gorm.io/gorm"
	"mini-tiktok-backend/pkg/consts"
	"strconv"
	"strings"
)

type Favorite struct {
	gorm.Model
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

type Video struct {
	gorm.Model
	AuthorId      int64  `json:"author_id"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	Title         string `json:"title"`
	FavoriteCount uint   `json:"favorite_count"`
	CommentCount  uint   `json:"comment_count"`
}

func (f *Favorite) TableName() string {
	return consts.FavoriteTableName
}

// GetVideoKey Key format is "video:{video_id}"
func GetVideoKey(videoId int64) string {
	var res strings.Builder
	res.WriteString("video:")
	res.WriteString(strconv.FormatInt(videoId, 10))
	return res.String()
}

func CreateFavorite(ctx context.Context, favorite *Favorite) error {
	var err error
	db := DB.Begin()

	if err = db.WithContext(ctx).Create(favorite).Error; err != nil {
		db.Rollback()
	}

	if err = db.WithContext(ctx).
		Model(&Video{}).
		Where("id = ?", favorite.VideoId).
		Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).
		Error; err != nil {
		db.Rollback()
	}

	db.Commit()

	// delete redis key for consistent
	RDB.HDel(ctx, GetVideoKey(favorite.VideoId), consts.FavoriteCount)

	return err
}

func DeleteFavorite(ctx context.Context, userId int64, videoId int64) error {
	var err error
	db := DB.Begin()

	if err = db.WithContext(ctx).
		Where("user_id = ? and video_id = ? ", userId, videoId).
		Delete(&Favorite{}).Error; err != nil {
		db.Rollback()
	}

	if err = db.WithContext(ctx).
		Model(&Video{}).
		Where("id = ?", videoId).
		Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).
		Error; err != nil {
		db.Rollback()
	}

	db.Commit()

	// delete redis key for consistent
	RDB.HDel(ctx, GetVideoKey(videoId), consts.FavoriteCount)

	return err
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
