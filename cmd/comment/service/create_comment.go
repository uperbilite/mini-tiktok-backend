package service

import (
	"context"
	"mini-tiktok-backend/cmd/comment/dal/db"
	"mini-tiktok-backend/cmd/comment/pack"
	"mini-tiktok-backend/cmd/comment/rpc"
	"mini-tiktok-backend/kitex_gen/comment"
	"mini-tiktok-backend/kitex_gen/user"
	"time"
)

type CreateCommentService struct {
	ctx context.Context
}

func NewCreateCommentService(ctx context.Context) *CreateCommentService {
	return &CreateCommentService{ctx: ctx}
}

func (s *CreateCommentService) CreateComment(req *comment.CreateCommentRequest) (*comment.Comment, error) {
	c, err := db.CreateComment(s.ctx, &db.Comment{
		UserId:     req.UserId,
		VideoId:    req.VideoId,
		Content:    req.Content,
		CreateDate: time.Now().Format("01-02"),
	})
	if err != nil {
		return nil, err
	}

	res := pack.Comment(c)

	u, err := rpc.QueryUser(s.ctx, &user.QueryUserRequest{
		UserId:       req.UserId,
		TargetUserId: req.UserId, // Current user creates this comment.
	})
	if err != nil {
		return nil, err
	}

	res.User = pack.User(u)

	return res, nil
}
