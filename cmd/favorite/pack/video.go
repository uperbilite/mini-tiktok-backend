package pack

import (
	"mini-tiktok-backend/kitex_gen/favorite"
	video2 "mini-tiktok-backend/kitex_gen/video"
)

func Video(v *video2.Video) *favorite.Video {
	if v == nil {
		return nil
	}
	return &favorite.Video{
		Id:            v.Id,
		Author:        User(v.Author),
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    v.IsFavorite,
		Title:         v.Title,
	}
}

func Videos(vs []*video2.Video) []*favorite.Video {
	videos := make([]*favorite.Video, 0)
	for _, v := range vs {
		if video := Video(v); video != nil {
			videos = append(videos, video)
		}
	}
	return videos
}
