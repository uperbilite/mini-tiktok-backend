package service

import (
	"context"
	"mini-tiktok-backend/cmd/favorite/dal/db"
	"mini-tiktok-backend/cmd/favorite/dal/mq"
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
		err = db.CreateFavorite(s.ctx, &db.Favorite{
			UserId:  req.UserId,
			VideoId: req.VideoId,
		})
		msg := &mq.Message{
			ActionType: 1,
			UserId:     req.UserId,
			VideoId:    req.VideoId,
		}
		msg.Produce()
		return err
	}
	if req.ActionType == 2 {
		return db.DeleteFavorite(s.ctx, req.UserId, req.VideoId)
	}

	return errno.ParamErr
}
