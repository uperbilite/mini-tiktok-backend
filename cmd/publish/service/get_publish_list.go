package service

import (
	"context"
	"mini-tiktok-backend/cmd/publish/dal/db"
	"mini-tiktok-backend/cmd/publish/pack"
	"mini-tiktok-backend/kitex_gen/publish"
)

type GetPublishListService struct {
	ctx context.Context
}

func NewGetPublishListService(ctx context.Context) *GetPublishListService {
	return &GetPublishListService{ctx: ctx}
}

func (s *GetPublishListService) GetPublishList(req *publish.GetPublishListRequest) ([]*publish.Video, error) {
	// TODO: get video list from id
	videos, err := db.MGetVideo(s.ctx, req.TargetUserId)
	if err != nil {
		return nil, err
	}
	// TODO: get user info from video author id
	// TODO: get favourite status of each video
	return pack.Videos(videos), nil
}
