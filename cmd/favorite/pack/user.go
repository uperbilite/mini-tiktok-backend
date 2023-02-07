package pack

import (
	"mini-tiktok-backend/kitex_gen/favorite"
	"mini-tiktok-backend/kitex_gen/video"
)

func User(u *video.User) *favorite.User {
	if u == nil {
		return nil
	}
	return &favorite.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
}
