package service

import (
	"context"
	"mini-tiktok-backend/cmd/relation/dal/db"
	"mini-tiktok-backend/cmd/relation/pack"
	"mini-tiktok-backend/cmd/relation/rpc"
	"mini-tiktok-backend/kitex_gen/relation"
	"mini-tiktok-backend/kitex_gen/user"
)

type GetFollowerListService struct {
	ctx context.Context
}

//NewGetFollowerListService new GetFollowerListService
func NewGetFollowerListService(ctx context.Context) *GetFollowerListService {
	return &GetFollowerListService{ctx}
}

//GetFollowerList get follower list
func (s *GetFollowerListService) GetFollowerList(req *relation.GetFollowerListRequest) ([]*relation.User,error) {
	userId := req.GetTargetUserId()
	followerId,err := db.QueryFollowerById(s.ctx,userId)
	if err != nil {
		return nil, err
	}
	var users []*relation.User
	for _, fd := range followerId {
		u, err := rpc.QueryUser(s.ctx,&user.QueryUserRequest{
			UserId: req.UserId,
			TargetUserId: fd.FromId,
		})
		if err != nil {
			return nil, err
		}
		r := pack.User(u)
		users = append(users,r)
	}
	return users,nil
}


