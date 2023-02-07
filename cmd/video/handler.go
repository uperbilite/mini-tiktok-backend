package main

import (
	"context"
	"mini-tiktok-backend/cmd/video/pack"
	"mini-tiktok-backend/cmd/video/service"
	video "mini-tiktok-backend/kitex_gen/video"
	"mini-tiktok-backend/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// GetVideos implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideos(ctx context.Context, req *video.GetVideosRequest) (resp *video.GetVideosResponse, err error) {
	resp = new(video.GetVideosResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	videos, err := service.NewGetVideosService(ctx).GetVideos(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err) // pack err message
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Videos = videos

	return resp, nil
}
