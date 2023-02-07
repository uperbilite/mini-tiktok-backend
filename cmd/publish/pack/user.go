package pack

import (
	"mini-tiktok-backend/kitex_gen/publish"
	"mini-tiktok-backend/kitex_gen/video"
)

func User(u *video.User) *publish.User {
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
