package service

import (
	"context"
	"mini-tiktok-backend/cmd/relation/dal/db"
	"mini-tiktok-backend/cmd/relation/pack"
	"mini-tiktok-backend/cmd/relation/rpc"
	"mini-tiktok-backend/kitex_gen/relation"
	"mini-tiktok-backend/kitex_gen/user"
)

type GetFriendListService struct {
	ctx context.Context
}

//NewFriendListService new GetFollowListService
func NewFriendListService(ctx context.Context) *GetFriendListService {
	return &GetFriendListService{ctx}
}

// GetFollowList get follow list
func (s *GetFriendListService) GetFriendList(req *relation.GetFriendListRequest) ([]*relation.User, error) {
	userId := req.GetTargetUserId()
	friendsId,err := db.QueryFriendById(s.ctx,userId)
	if err != nil {
		return nil, err
	}
	var users []*relation.User
	for _, fd := range friendsId {
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
