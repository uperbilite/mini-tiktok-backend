package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"mini-tiktok-backend/kitex_gen/favorite"
	"mini-tiktok-backend/kitex_gen/favorite/favoriteservice"
	"mini-tiktok-backend/pkg/errno"
)

var favoriteClient favoriteservice.Client

func initFavorite() {
	c, err := favoriteservice.NewClient("favorite", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}
	favoriteClient = c
}

func GetIsFavorite(ctx context.Context, req *favorite.GetIsFavoriteRequest) (bool, error) {
	resp, err := favoriteClient.GetIsFavorite(ctx, req)
	if err != nil {
		return false, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return false, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.IsFavorite, nil
}
