package service

import (
	"context"
	"mini-tiktok-backend/cmd/comment/dal/db"
	"mini-tiktok-backend/kitex_gen/comment"
)

type DeleteCommentService struct {
	ctx context.Context
}

func NewDeleteCommentService(ctx context.Context) *DeleteCommentService {
	return &DeleteCommentService{ctx: ctx}
}

func (s *DeleteCommentService) DeleteService(req *comment.DeleteCommentRequest) error {
	return db.DeleteComment(s.ctx, req.CommentId)
}
