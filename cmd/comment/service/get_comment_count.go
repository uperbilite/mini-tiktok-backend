package service

import (
	"context"
	"mini-tiktok-backend/cmd/comment/dal/db"
	"mini-tiktok-backend/kitex_gen/comment"
)

type GetCommentCountService struct {
	ctx context.Context
}

func NewGetCommentCountService(ctx context.Context) *GetCommentCountService {
	return &GetCommentCountService{ctx: ctx}
}

func (s *GetCommentCountService) GetCommentCount(req *comment.GetCommentCountRequest) (int64, error) {
	commentCount, err := db.QueryCommentCount(s.ctx, req.VideoId)
	if err != nil {
		return 0, err
	}

	return commentCount, nil
}
