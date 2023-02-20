package service

import (
	"context"
	"mini-tiktok-backend/cmd/relation/dal/db"
	"mini-tiktok-backend/kitex_gen/relation"
)

type IsFollowToUserService struct {
	ctx context.Context
}

func NewIsFollowToUserService(ctx context.Context) *IsFollowToUserService {
	return &IsFollowToUserService{ctx}
}

func (s *IsFollowToUserService) IsFollowToUser(req *relation.IsFollowToUserRequest) (bool, error) {
	userId, toUserId := req.GetUserId(), req.GetToUserId()
	isFollow, err := db.IsFollow(s.ctx, userId, toUserId)
	if err != nil {
		return false, err
	}
	return isFollow, nil
}
