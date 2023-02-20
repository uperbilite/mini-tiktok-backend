package service

import (
	"context"
	"log"
	"mini-tiktok-backend/cmd/relation/dal/db"
	"mini-tiktok-backend/kitex_gen/relation"
)

type GetFollowAndFollowerCountService struct {
	ctx context.Context
}

func NewGetFollowAndFollowerCountService(ctx context.Context) *GetFollowAndFollowerCountService {
	return &GetFollowAndFollowerCountService{ctx}
}

func (s *GetFollowAndFollowerCountService) GetFollowAndFollowerCount(req *relation.GetFollowAndFollowerCountRequest) (follows, followers int64, err error) {
	userId := req.GetUserId()
	follows, followers, err = db.GetFollowNumAndFollowerNumFromRedis(s.ctx, userId)
	if err == nil {
		return
	}
	follows, err = db.CountFollow(s.ctx, userId)
	if err != nil {
		return 0, 0, err
	}
	followers, err = db.CountFollower(s.ctx, userId)
	if err != nil {
		return 0, 0, err
	}
	err = db.SetFollowNumAndFollowerNumToRedis(s.ctx, userId, follows, followers)
	if err != nil {
		log.Println("can't storage into redis")
		err = nil
	}
	return
}
