package pack

import (
	"mini-tiktok-backend/kitex_gen/relation"
	"mini-tiktok-backend/kitex_gen/user"
)

func User(u *user.User) *relation.User {
	if u == nil {
		return nil
	}
	return &relation.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
}
