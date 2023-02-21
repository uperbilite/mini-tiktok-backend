package service

import (
	"context"
	"log"
	"mini-tiktok-backend/cmd/relation/dal/db"
	"mini-tiktok-backend/kitex_gen/relation"
)

type RelationActionService struct {
	ctx context.Context
}

func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

func (s *RelationActionService) RelationAction(req *relation.RelationActionRequest) (err error) {
	fromId, toId, actionType := req.GetUserId(), req.GetToUserId(), req.GetActionType()
	if actionType == 1 {
		if err = db.FollowUser(s.ctx, fromId, toId); err != nil {
			return
		}
	} else if actionType == 2 {
		if err = db.CancelFollow(s.ctx, fromId, toId); err != nil {
			return
		}
	}
	err = db.RemoveKeyFromRedis(s.ctx,fromId,toId)
	if err != nil {
		log.Println("redis err")
		err = nil
	}
	return
}
