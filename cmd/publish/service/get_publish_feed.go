package service

import (
	"context"
	"mini-tiktok-backend/cmd/publish/dal/db"
	"mini-tiktok-backend/cmd/publish/pack"
	"mini-tiktok-backend/cmd/publish/rpc"
	"mini-tiktok-backend/kitex_gen/publish"
	"mini-tiktok-backend/kitex_gen/video"
	"time"
)

type GetPublishFeedService struct {
	ctx context.Context
}

func NewGetPublishFeedService(ctx context.Context) *GetPublishFeedService {
	return &GetPublishFeedService{ctx: ctx}
}

func (s *GetPublishFeedService) GetPublishFeed(req *publish.GetPublishFeedRequest) ([]*publish.Video, int64, error) {
	videoIds, err := db.GetVideoIdsFeed(s.ctx, req.LatestTime)
	if err != nil {
		return nil, 0, err
	}

	if len(videoIds) == 0 {
		// No video feed will return current timestamp.
		return make([]*publish.Video, 0), time.Now().UnixMilli(), nil
	}

	videos, err := rpc.GetVideos(s.ctx, &video.GetVideosRequest{
		UserId:   req.UserId,
		VideoIds: videoIds,
	})
	if err != nil {
		return nil, 0, err
	}

	return pack.Videos(videos), time.Now().UnixMilli(), nil // TODO: get next_time from latest video create time
}
