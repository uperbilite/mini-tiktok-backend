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

type GetPublishFeedService struct {
	ctx context.Context
}

func NewGetPublishFeedService(ctx context.Context) *GetPublishFeedService {
	return &GetPublishFeedService{ctx: ctx}
}

func (s *GetPublishFeedService) GetPublishFeed(req *publish.GetPublishFeedRequest) ([]*publish.Video, int64, error) {
	vs, err := db.GetVideoFeed(s.ctx, req.LatestTime)
	if err != nil {
		return nil, 0, err
	}

	if len(vs) == 0 {
		return nil, 0, errno.AuthorizationFailedErr // TODO: set errno msg
	}

	videos := make([]*publish.Video, 0)

	// TODO: get user info from video author id, using UserId and TargetUserId
	for _, v := range vs {
		video := pack.Video(v)
		resp, _ := rpc.QueryUser(s.ctx, &user.QueryUserRequest{
			UserId:       req.UserId,
			TargetUserId: v.AuthorId,
		})
		video.Author = pack.User(resp)
		videos = append(videos, video)
		// TODO: get favourite status of each video
	}

	return videos, vs[0].Model.CreatedAt.UnixMilli(), nil
}
