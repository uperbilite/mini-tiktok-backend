package pack

import (
	"mini-tiktok-backend/kitex_gen/comment"
	"mini-tiktok-backend/kitex_gen/user"
)

func User(u *user.User) *comment.User {
	if u == nil {
		return nil
	}
	return &comment.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
}
