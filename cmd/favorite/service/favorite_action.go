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
		// TODO: favorite exists error
		err := db.CreateFavoriteInRedis(s.ctx, req.VideoId)
		msg := &mq.Message{
			ActionType: 1,
			UserId:     req.UserId,
			VideoId:    req.VideoId,
		}
		msg.Produce()
		return err
	}
	if req.ActionType == 2 {
		// TODO: favorite not exists error
		err := db.DeleteFavoriteInRedis(s.ctx, req.VideoId)
		msg := &mq.Message{
			ActionType: 2,
			UserId:     req.UserId,
			VideoId:    req.VideoId,
		}
		msg.Produce()
		return err
	}

	return errno.ParamErr
}
