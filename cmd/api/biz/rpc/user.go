package rpc

import (
	"context"
	client2 "github.com/cloudwego/kitex/client"
	"mini-tiktok-backend/kitex_gen/user"
	"mini-tiktok-backend/kitex_gen/user/userservice"
)

var userClient userservice.Client

func initUser() {
	// TODO: get user service from etcd
	client, err := userservice.NewClient("user", client2.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}
	userClient = client
}

// CheckUser check user info
func CheckUser(ctx context.Context, req *user.CheckUserRequest) (int64, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, err // TODO: create new errno
	}
	return resp.UserId, nil
}
