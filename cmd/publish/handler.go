package main

import (
	"context"
	publish "mini-tiktok-backend/kitex_gen/publish"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishVideo implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishVideo(ctx context.Context, req *publish.PublishVideoRequest) (resp *publish.PublishVideoResponse, err error) {
	resp = new(publish.PublishVideoResponse)
	resp.BaseResp = new(publish.BaseResp)
	return resp, nil
}
