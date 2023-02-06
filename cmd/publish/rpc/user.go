package rpc

import (
	"context"
	client2 "github.com/cloudwego/kitex/client"
	"mini-tiktok-backend/kitex_gen/user"
	"mini-tiktok-backend/kitex_gen/user/userservice"
	"mini-tiktok-backend/pkg/errno"
)

var userClient userservice.Client

func initUser() {
	client, err := userservice.NewClient("user", client2.WithHostPorts("127.0.0.1:8081"))
	if err != nil {
		panic(err)
	}
	userClient = client
}

func QueryUser(ctx context.Context, req *user.QueryUserRequest) (*user.User, error) {
	resp, err := userClient.QueryUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.User, nil
}
