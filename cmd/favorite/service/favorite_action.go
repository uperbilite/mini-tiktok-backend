package service

import (
	"context"
	"mini-tiktok-backend/cmd/favorite/dal/db"
	"mini-tiktok-backend/cmd/favorite/dal/mq"
	"mini-tiktok-backend/kitex_gen/favorite"
	"mini-tiktok-backend/pkg/consts"
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
		lock := db.NewFavoriteKeyLock(req.VideoId, consts.FavoriteCount)
		db.CreateFavoriteInRedis(s.ctx, req.VideoId)
		lock.Unlock(s.ctx)
		msg := &mq.Message{
			ActionType: 1,
			UserId:     req.UserId,
			VideoId:    req.VideoId,
		}
		// mq for create favorite in
		msg.Produce()
		return nil
	}
	if req.ActionType == 2 {
		// TODO: favorite not exists error
		lock := db.NewFavoriteKeyLock(req.VideoId, consts.FavoriteCount)
		db.DeleteFavoriteInRedis(s.ctx, req.VideoId)
		lock.Unlock(s.ctx)
		msg := &mq.Message{
			ActionType: 2,
			UserId:     req.UserId,
			VideoId:    req.VideoId,
		}
		msg.Produce()
		return nil
	}

	return errno.ParamErr
}
