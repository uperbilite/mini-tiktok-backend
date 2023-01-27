package main

import (
	"context"
	"mini-tiktok-backend/cmd/user/pack"
	"mini-tiktok-backend/cmd/user/service"
	user "mini-tiktok-backend/kitex_gen/user"
	"mini-tiktok-backend/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	resp = new(user.CheckUserResponse)

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err) // pack err message
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserId = uid

	return resp, nil
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	resp = new(user.CreateUserResponse)

	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err) // pack err message
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)

	return resp, nil
}
