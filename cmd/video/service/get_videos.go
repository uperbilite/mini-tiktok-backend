package service

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mini-tiktok-backend/cmd/video/dal/db"
	"mini-tiktok-backend/cmd/video/pack"
	"mini-tiktok-backend/cmd/video/rpc"
	"mini-tiktok-backend/kitex_gen/favorite"
	"mini-tiktok-backend/kitex_gen/user"
	video2 "mini-tiktok-backend/kitex_gen/video"
	"mini-tiktok-backend/pkg/consts"
)

type GetVideosService struct {
	ctx context.Context
}

func NewGetVideosService(ctx context.Context) *GetVideosService {
	return &GetVideosService{ctx: ctx}
}

func (s *GetVideosService) GetVideos(req *video2.GetVideosRequest) ([]*video2.Video, error) {
	vs, err := db.MGetVideos(s.ctx, req.VideoIds)
	if err != nil {
		return nil, err
	}

	client, err := oss.New(consts.OSSEndPoint, consts.AccessKeyId, consts.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(consts.OSSBucketName)
	if err != nil {
		return nil, err
	}

	videos := make([]*video2.Video, 0)

	for _, v := range vs {
		video := pack.Video(v)
		if video == nil {
			continue
		}

		videoSignedUrl, err := bucket.SignURL(v.PlayURL, oss.HTTPGet, 600)
		if err != nil {
			return nil, err
		}
		coverSignedUrl, err := bucket.SignURL(v.CoverURL, oss.HTTPGet, 600)
		if err != nil {
			return nil, err
		}
		video.PlayUrl = videoSignedUrl
		video.CoverUrl = coverSignedUrl

		// get user info
		resp, err := rpc.QueryUser(s.ctx, &user.QueryUserRequest{
			UserId:       req.UserId,
			TargetUserId: v.AuthorId,
		})
		if err != nil {
			return nil, err
		}
		video.Author = pack.User(resp)

		// get is_favorite
		isFavorite, err := rpc.GetIsFavorite(s.ctx, &favorite.GetIsFavoriteRequest{
			UserId:  req.UserId,
			VideoId: int64(v.ID),
		})
		if err != nil {
			return nil, err
		}
		video.IsFavorite = isFavorite

		video.FavoriteCount, _ = db.GetFavoriteCount(s.ctx, int64(v.ID))
		video.CommentCount, _ = db.GetCommentCount(s.ctx, int64(v.ID))

		videos = append(videos, video)
	}

	return videos, nil
}
