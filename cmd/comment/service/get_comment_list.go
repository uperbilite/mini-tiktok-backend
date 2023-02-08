package service

import (
	"context"
	"mini-tiktok-backend/cmd/comment/dal/db"
	"mini-tiktok-backend/cmd/comment/pack"
	"mini-tiktok-backend/cmd/comment/rpc"
	"mini-tiktok-backend/kitex_gen/comment"
	"mini-tiktok-backend/kitex_gen/user"
)

type GetCommentListService struct {
	ctx context.Context
}

func NewGetCommentListService(ctx context.Context) *GetCommentListService {
	return &GetCommentListService{ctx: ctx}
}

func (s *GetCommentListService) GetCommentList(req *comment.GetCommentListRequest) ([]*comment.Comment, error) {
	cs, err := db.GetCommentsByVideoId(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	comments := make([]*comment.Comment, 0)

	for _, c := range cs {
		m := pack.Comment(c)

		u, err := rpc.QueryUser(s.ctx, &user.QueryUserRequest{
			UserId:       req.UserId,
			TargetUserId: c.UserId, // The relation between current user and comment user
		})
		if err != nil {
			return nil, err
		}
		m.User = pack.User(u)

		comments = append(comments, m)
	}

	return comments, nil
}
