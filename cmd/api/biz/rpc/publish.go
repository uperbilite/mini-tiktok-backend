package rpc

import (
	"context"
	client2 "github.com/cloudwego/kitex/client"
	"mini-tiktok-backend/kitex_gen/publish"
	"mini-tiktok-backend/kitex_gen/publish/publishservice"
	"mini-tiktok-backend/pkg/errno"
)

var publishClient publishservice.Client

func initPublish() {
	client, err := publishservice.NewClient("publish", client2.WithHostPorts("127.0.0.1:8082"))
	if err != nil {
		panic(err)
	}
	publishClient = client
}

func PublishVideo(ctx context.Context, req *publish.PublishVideoRequest) error {
	resp, err := publishClient.PublishVideo(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 { // unpack err message from resp
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return err
}

func GetPublishList(ctx context.Context, req *publish.GetPublishListRequest) ([]*publish.Video, error) {
	resp, err := publishClient.GetPublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 { // unpack err message from resp
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.VideoList, nil
}

func GetPublishFeed(ctx context.Context, req *publish.GetPublishFeedRequest) ([]*publish.Video, int64, error) {
	resp, err := publishClient.GetPublishFeed(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	if resp.BaseResp.StatusCode != 0 { // unpack err message from resp
		return nil, 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.VideoList, resp.NextTime, nil
}
