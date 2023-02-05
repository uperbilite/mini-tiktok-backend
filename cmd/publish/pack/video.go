package pack

import (
	"mini-tiktok-backend/cmd/publish/dal/db"
	"mini-tiktok-backend/kitex_gen/publish"
)

func Video(v *db.Video) *publish.Video {
	if v == nil {
		return nil
	}
	return &publish.Video{
		Id: int64(v.ID),
		Author: &publish.User{
			Id:            1,
			Name:          "123123",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}, // TODO: Get author info from user service
		PlayUrl:       v.PlayURL,
		CoverUrl:      v.CoverURL,
		FavoriteCount: int64(v.FavouriteCount),
		CommentCount:  int64(v.CommentCount),
		IsFavorite:    false, // TODO: Get is favourite from favourite service
		Title:         v.Title,
	}
}

func Videos(vs []*db.Video) []*publish.Video {
	videos := make([]*publish.Video, 0)
	for _, v := range vs {
		if temp := Video(v); temp != nil {
			videos = append(videos, temp)
		}
	}
	return videos
}
