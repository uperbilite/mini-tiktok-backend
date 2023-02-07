package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"mini-tiktok-backend/kitex_gen/video"
	"mini-tiktok-backend/kitex_gen/video/videoservice"
	"mini-tiktok-backend/pkg/errno"
)

var videoClient videoservice.Client

func initVideo() {
	c, err := videoservice.NewClient("video", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func GetVideos(ctx context.Context, req *video.GetVideosRequest) ([]*video.Video, error) {
	resp, err := videoClient.GetVideos(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.Videos, nil
}
