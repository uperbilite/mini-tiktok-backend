package pack

import (
	"mini-tiktok-backend/cmd/publish/dal/db"
	"mini-tiktok-backend/kitex_gen/publish"
	"mini-tiktok-backend/kitex_gen/user"
)

func Video(v *db.Video) *publish.Video {
	if v == nil {
		return nil
	}
	return &publish.Video{
		Id:            int64(v.ID),
		Author:        nil, // TODO: Get author info from user service
		PlayUrl:       v.PlayURL,
		CoverUrl:      v.CoverURL,
		FavoriteCount: int64(v.FavouriteCount),
		CommentCount:  int64(v.CommentCount),
		IsFavorite:    false, // TODO: Get is favourite from favourite service
		Title:         v.Title,
	}
}

func User(u *user.User) *publish.User {
	if u == nil {
		return nil
	}
	return &publish.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
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
