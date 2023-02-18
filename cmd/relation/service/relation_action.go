package service

import (
	"context"
	"mini-tiktok-backend/cmd/relation/dal/db"
	"mini-tiktok-backend/kitex_gen/relation"
)

type RelationActionService struct {
	ctx context.Context
}

func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

func (s *RelationActionService) RelationAction(req *relation.RelationActionRequest) error {
	userId,targetUserId := req.GetUserId(),req.GetToUserId()
	if err := db.FollowUser(s.ctx,userId,targetUserId); err != nil {
		return err
	}
	return nil
}