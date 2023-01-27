package service

import (
	"context"
	"mini-tiktok-backend/cmd/user/dal/db"
	"mini-tiktok-backend/kitex_gen/user"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *user.CreateUserRequest) error {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return err
	}

	return db.CreateUser(s.ctx, []*db.User{{
		Username: req.Username,
		Password: req.Password,
	}})
}
