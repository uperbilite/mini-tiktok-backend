package service

import (
	"context"
	"mini-tiktok-backend/cmd/user/dal/db"
	"mini-tiktok-backend/pkg/errno"

	"mini-tiktok-backend/kitex_gen/user"
)

type CheckUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

// CheckUser check user info
func (s *CheckUserService) CheckUser(req *user.CheckUserRequest) (int64, error) {
	// TODO: use encrypted password and then compare username and password

	userName := req.Username
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return 0, err
	}

	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}

	id := users[0].ID
	return int64(id), nil
}
