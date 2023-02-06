package service

import (
	"context"
	"mini-tiktok-backend/cmd/publish/dal/db"
	"mini-tiktok-backend/cmd/publish/pack"
	"mini-tiktok-backend/cmd/publish/rpc"
	"mini-tiktok-backend/kitex_gen/publish"
	"mini-tiktok-backend/kitex_gen/user"
	"mini-tiktok-backend/pkg/errno"
)

type GetPublishListService struct {
	ctx context.Context
}

func NewGetPublishListService(ctx context.Context) *GetPublishListService {
	return &GetPublishListService{ctx: ctx}
}

func (s *GetPublishListService) GetPublishList(req *publish.GetPublishListRequest) ([]*publish.Video, error) {
	// TODO: get video list from id
	vs, err := db.MGetVideo(s.ctx, req.TargetUserId)
	if err != nil {
		return nil, err
	}

	if len(vs) == 0 {
		return nil, errno.AuthorizationFailedErr // TODO: set error msg
	}

	// TODO: get user info from video author id, using UserId and TargetUserId
	videos := make([]*publish.Video, 0)

	for _, v := range vs {
		video := pack.Video(v)
		resp, err := rpc.QueryUser(s.ctx, &user.QueryUserRequest{
			UserId:       req.UserId,
			TargetUserId: req.TargetUserId,
		})
		if err != nil {
			return nil, err
		}
		video.Author = pack.User(resp)
		videos = append(videos, video)
		// TODO: get favourite status of each video
	}

	return videos, nil
}
