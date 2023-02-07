package pack

import (
	"mini-tiktok-backend/kitex_gen/publish"
	video2 "mini-tiktok-backend/kitex_gen/video"
)

func Video(v *video2.Video) *publish.Video {
	if v == nil {
		return nil
	}
	return &publish.Video{
		Id:            v.Id,
		Author:        User(v.Author), // TODO: Get author info from user service
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount, // TODO: Get favorite count from favorite service
		CommentCount:  v.CommentCount,  // TODO: Get comment count from comment service
		IsFavorite:    v.IsFavorite,    // TODO: Get is favourite from favorite service
		Title:         v.Title,
	}
}

func Videos(vs []*video2.Video) []*publish.Video {
	videos := make([]*publish.Video, 0)
	for _, v := range vs {
		if video := Video(v); video != nil {
			videos = append(videos, video)
		}
	}
	return videos
}
