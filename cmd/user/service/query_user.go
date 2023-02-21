package service

import (
	"context"
	"mini-tiktok-backend/cmd/user/dal/db"
	"mini-tiktok-backend/cmd/user/pack"
	"mini-tiktok-backend/cmd/user/rpc"
	"mini-tiktok-backend/kitex_gen/relation"
	"mini-tiktok-backend/kitex_gen/user"
	"mini-tiktok-backend/pkg/errno"
)

type QueryUserService struct {
	ctx context.Context
}

func NewQueryUserService(ctx context.Context) *QueryUserService {
	return &QueryUserService{ctx: ctx}
}

func (s *QueryUserService) QueryUser(req *user.QueryUserRequest) (*user.User, error) {
	users, err := db.QueryUserById(s.ctx, req.TargetUserId)

	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.UserNotExistErr
	}
	r := pack.User(users[0])
	follows, followers, err := rpc.GetFollowAndFollowerCount(s.ctx, &relation.GetFollowAndFollowerCountRequest{
		UserId: req.TargetUserId,
	})
	if err != nil {
		return nil, err
	}
	r.FollowCount = follows
	r.FollowerCount = followers
	if req.UserId != 0 {
		isFollow, err := rpc.IsFollowToUser(s.ctx, &relation.IsFollowToUserRequest{
			UserId:   req.GetUserId(),
			ToUserId: req.GetTargetUserId(),
		})
		if err != nil {
			return nil, err
		}
		r.IsFollow = isFollow
	}
	return r, nil
}
