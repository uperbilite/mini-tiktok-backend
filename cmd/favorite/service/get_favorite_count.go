package service

import (
	"context"
	"mini-tiktok-backend/cmd/favorite/dal/db"
	"mini-tiktok-backend/kitex_gen/favorite"
)

type GetFavoriteCountService struct {
	ctx context.Context
}

func NewGetFavoriteCountService(ctx context.Context) *GetFavoriteCountService {
	return &GetFavoriteCountService{
		ctx: ctx,
	}
}

func (s *GetFavoriteCountService) GetFavoriteCount(req *favorite.GetFavoriteCountRequest) (int64, error) {
	favoriteCount, err := db.QueryFavoriteCount(s.ctx, req.VideoId)
	if err != nil {
		return 0, err
	}

	return favoriteCount, nil
}
