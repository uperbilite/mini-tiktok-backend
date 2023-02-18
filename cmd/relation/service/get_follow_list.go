package service

import (
	"context"
	"mini-tiktok-backend/cmd/relation/dal/db"
	"mini-tiktok-backend/cmd/relation/pack"
	"mini-tiktok-backend/cmd/relation/rpc"
	"mini-tiktok-backend/kitex_gen/relation"
	"mini-tiktok-backend/kitex_gen/user"
)

type GetFollowListService struct {
	ctx context.Context
}

//NewGetFollowListService new GetFollowListService
func NewGetFollowListService(ctx context.Context) *GetFollowListService {
	return &GetFollowListService{ctx}
}

// GetFollowList get follow list
func (s *GetFollowListService) GetFollowList(req *relation.GetFollowListRequest) ([]*relation.User, error) {
	userId := req.GetTargetUserId()
	followsId,err := db.QueryFollowById(s.ctx,userId)
	if err != nil {
		return nil, err
	}
	var users []*relation.User
	for _, fd := range followsId {
		u, err := rpc.QueryUser(s.ctx,&user.QueryUserRequest{
			UserId: req.UserId,
			TargetUserId: fd.FollowId,
		})
		if err != nil {
			return nil, err
		}
		r := pack.User(u)
		users = append(users,r)
	}
	return users,nil
}
