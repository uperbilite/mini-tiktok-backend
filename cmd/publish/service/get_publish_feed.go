package service

import (
	"context"
	"mini-tiktok-backend/cmd/publish/dal/db"
	"mini-tiktok-backend/cmd/publish/pack"
	"mini-tiktok-backend/cmd/publish/rpc"
	"mini-tiktok-backend/kitex_gen/publish"
	"mini-tiktok-backend/kitex_gen/user"
	"time"
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

	videos := make([]*publish.Video, 0)

	if len(vs) == 0 {
		// No video feed will return current timestamp.
		return videos, time.Now().UnixMilli(), nil
	}

	// TODO: get user info from video author id, using UserId and TargetUserId
	for _, v := range vs {
		video := pack.Video(v)
		resp, _ := rpc.QueryUser(s.ctx, &user.QueryUserRequest{
			UserId:       req.UserId,
			TargetUserId: v.AuthorId,
		})
		// TODO: error handle
		video.Author = pack.User(resp)
		videos = append(videos, video)
		// TODO: get favourite status of each video
	}

	return videos, vs[0].Model.CreatedAt.UnixMilli(), nil
}
