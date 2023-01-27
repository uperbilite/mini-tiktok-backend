package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
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
	username := req.Username
	users, err := db.QueryUser(s.ctx, username)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}

	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	u := users[0]
	if u.Password != password {
		return 0, errno.AuthorizationFailedErr
	}

	id := u.ID
	return int64(id), nil
}
