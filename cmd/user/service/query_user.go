package service

import (
	"context"
	"mini-tiktok-backend/cmd/user/dal/db"
	"mini-tiktok-backend/kitex_gen/user"
	"mini-tiktok-backend/pkg/errno"
)

type QueryUserService struct {
	ctx context.Context
}

func NewQueryUserService(ctx context.Context) *QueryUserService {
	return &QueryUserService{
		ctx: ctx,
	}
}

func (s *QueryUserService) QueryUser(req *user.QueryUserRequest) (string, error) {
	id := req.UserId
	users, err := db.QueryUserById(s.ctx, id)

	if err != nil {
		return "", err
	}
	if len(users) == 0 {
		return "", errno.AuthorizationFailedErr // TODO: id not exist error
	}

	u := users[0]
	username := u.Username

	return username, nil
}
