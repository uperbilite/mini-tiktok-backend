package service

import (
	"context"
	"mini-tiktok-backend/cmd/favorite/dal/db"
	"mini-tiktok-backend/cmd/favorite/pack"
	"mini-tiktok-backend/cmd/favorite/rpc"
	"mini-tiktok-backend/kitex_gen/favorite"
	"mini-tiktok-backend/kitex_gen/video"
)

type GetFavoriteListService struct {
	ctx context.Context
}

func NewGetFavoriteListService(ctx context.Context) *GetFavoriteListService {
	return &GetFavoriteListService{ctx: ctx}
}

func (s *GetFavoriteListService) GetFavoriteList(req *favorite.GetFavoriteListRequest) ([]*favorite.Video, error) {
	videoIds, err := db.GetVideoIdsByUserId(s.ctx, req.TargetUserId)
	if err != nil {
		return nil, err
	}

	videos, err := rpc.GetVideos(s.ctx, &video.GetVideosRequest{
		UserId:   req.UserId,
		VideoIds: videoIds,
	})
	if err != nil {
		return nil, err
	}

	return pack.Videos(videos), nil
}
