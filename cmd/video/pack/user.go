package pack

import (
	"mini-tiktok-backend/kitex_gen/user"
	"mini-tiktok-backend/kitex_gen/video"
)

func User(u *user.User) *video.User {
	if u == nil {
		return nil
	}
	return &video.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
}
