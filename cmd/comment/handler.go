package main

import (
	"context"
	"mini-tiktok-backend/cmd/comment/pack"
	"mini-tiktok-backend/cmd/comment/service"
	comment "mini-tiktok-backend/kitex_gen/comment"
	"mini-tiktok-backend/pkg/errno"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CreateComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CreateComment(ctx context.Context, req *comment.CreateCommentRequest) (resp *comment.CreateCommentResponse, err error) {
	resp = new(comment.CreateCommentResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	c, err := service.NewCreateCommentService(ctx).CreateComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err) // pack err message
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Comment = c

	return resp, nil
}

// DeleteComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) DeleteComment(ctx context.Context, req *comment.DeleteCommentRequest) (resp *comment.DeleteCommentResponse, err error) {
	// TODO: Your code here...
	return
}
