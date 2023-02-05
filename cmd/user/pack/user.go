package pack

import (
	"mini-tiktok-backend/cmd/user/dal/db"
	"mini-tiktok-backend/kitex_gen/user"
)

func User(u *db.User) *user.User {
	if u == nil {
		return nil
	}
	return &user.User{
		Id:            int64(u.ID),
		Name:          u.Username,
		FollowCount:   0, // TODO: Get these three params from relation service
		FollowerCount: 0,
		IsFollow:      false,
	}
}
