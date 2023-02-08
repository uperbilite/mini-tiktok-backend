package service

import (
	"context"
	"mini-tiktok-backend/cmd/user/dal/db"
	"mini-tiktok-backend/cmd/user/pack"
	"mini-tiktok-backend/kitex_gen/user"
	"mini-tiktok-backend/pkg/errno"
)

type QueryUserService struct {
	ctx context.Context
}

func NewQueryUserService(ctx context.Context) *QueryUserService {
	return &QueryUserService{ctx: ctx}
}

func (s *QueryUserService) QueryUser(req *user.QueryUserRequest) (*user.User, error) {
	users, err := db.QueryUserById(s.ctx, req.TargetUserId)

	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.UserNotExistErr
	}

	if req.UserId == 0 {
		// TODO: Get follow and follower count
	} else {
		// TODO: Get follow and follower count and is_followed
	}

	return pack.User(users[0]), nil
}
