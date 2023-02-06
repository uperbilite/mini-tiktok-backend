package main

import (
	"context"
	"mini-tiktok-backend/cmd/publish/pack"
	"mini-tiktok-backend/cmd/publish/service"
	publish "mini-tiktok-backend/kitex_gen/publish"
	"mini-tiktok-backend/pkg/errno"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishVideo implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishVideo(ctx context.Context, req *publish.PublishVideoRequest) (resp *publish.PublishVideoResponse, err error) {
	resp = new(publish.PublishVideoResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewPublishVideoService(ctx).PublishVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err) // pack err message
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)

	return resp, nil
}

func (s *PublishServiceImpl) GetPublishList(ctx context.Context, req *publish.GetPublishListRequest) (resp *publish.GetPublishListResponse, err error) {
	resp = new(publish.GetPublishListResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	videoList, err := service.NewGetPublishListService(ctx).GetPublishList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videoList

	return resp, nil
}

func (s *PublishServiceImpl) GetPublishFeed(ctx context.Context, req *publish.GetPublishFeedRequest) (resp *publish.GetPublishFeedResponse, err error) {
	resp = new(publish.GetPublishFeedResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	videoList, nextTime, err := service.NewGetPublishFeedService(ctx).GetPublishFeed(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videoList
	resp.NextTime = nextTime

	return resp, nil
}
