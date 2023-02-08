package service

import (
	"context"
	"mini-tiktok-backend/cmd/favorite/dal/db"
	"mini-tiktok-backend/kitex_gen/favorite"
	"mini-tiktok-backend/pkg/errno"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{
		ctx: ctx,
	}
}

func (s *FavoriteActionService) FavoriteAction(req *favorite.FavoriteActionRequest) error {
	if req.ActionType == 1 {
		favorites, err := db.QueryFavorite(s.ctx, req.UserId, req.VideoId)
		if err != nil {
			return err
		}
		if len(favorites) != 0 {
			return errno.UserAlreadyExistErr // TODO: FavoriteAlreadyExistErr
		}
		return db.CreateFavorite(s.ctx, &db.Favorite{
			UserId:  req.UserId,
			VideoId: req.VideoId,
		})
	}
	if req.ActionType == 2 {
		return db.DeleteFavorite(s.ctx, req.UserId, req.VideoId)
	}

	return errno.ParamErr
}
