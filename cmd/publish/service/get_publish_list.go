package service

import (
	"context"
	"mini-tiktok-backend/cmd/publish/dal/db"
	"mini-tiktok-backend/cmd/publish/pack"
	"mini-tiktok-backend/cmd/publish/rpc"
	"mini-tiktok-backend/kitex_gen/publish"
	"mini-tiktok-backend/kitex_gen/video"
)

type GetPublishListService struct {
	ctx context.Context
}

func NewGetPublishListService(ctx context.Context) *GetPublishListService {
	return &GetPublishListService{ctx: ctx}
}

func (s *GetPublishListService) GetPublishList(req *publish.GetPublishListRequest) ([]*publish.Video, error) {
	videoIds, err := db.GetVideoIdsByAuthorId(s.ctx, req.TargetUserId)
	if err != nil {
		return nil, err
	}

	videos, _ := rpc.GetVideos(s.ctx, &video.GetVideosRequest{
		UserId:   req.UserId,
		VideoIds: videoIds,
	})
	// TODO: error handle

	return pack.Videos(videos), nil
}
