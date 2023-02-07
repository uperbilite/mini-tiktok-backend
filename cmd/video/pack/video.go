package pack

import (
	"mini-tiktok-backend/cmd/video/dal/db"
	"mini-tiktok-backend/kitex_gen/video"
)

func Video(v *db.Video) *video.Video {
	if v == nil {
		return nil
	}
	return &video.Video{
		Id:            int64(v.ID),
		Author:        nil,
		PlayUrl:       v.PlayURL,
		CoverUrl:      v.CoverURL,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         v.Title,
	}
}
