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
	"sync"
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

	type item struct {
		video *video2.Video
		err   error
	}
	ch := make(chan item)
	var wg sync.WaitGroup

	for _, v := range vs {
		wg.Add(1)
		go func(v *db.Video) {
			defer wg.Done()

			var i item
			i.video = pack.Video(v)

			videoSignedUrl, err := bucket.SignURL(v.PlayURL, oss.HTTPGet, 600)
			if err != nil {
				i.err = err
				ch <- i
				return
			}
			coverSignedUrl, err := bucket.SignURL(v.CoverURL, oss.HTTPGet, 600)
			if err != nil {
				i.err = err
				ch <- i
				return
			}
			i.video.PlayUrl = videoSignedUrl
			i.video.CoverUrl = coverSignedUrl

			// get user info
			resp, err := rpc.QueryUser(s.ctx, &user.QueryUserRequest{
				UserId:       req.UserId,
				TargetUserId: v.AuthorId,
			})
			if err != nil {
				i.err = err
				ch <- i
				return
			}
			i.video.Author = pack.User(resp)

			// get is_favorite
			isFavorite, err := rpc.GetIsFavorite(s.ctx, &favorite.GetIsFavoriteRequest{
				UserId:  req.UserId,
				VideoId: int64(v.ID),
			})
			if err != nil {
				i.err = err
				ch <- i
				return
			}
			i.video.IsFavorite = isFavorite

			i.video.FavoriteCount, _ = db.GetFavoriteCount(s.ctx, int64(v.ID))
			i.video.CommentCount, _ = db.GetCommentCount(s.ctx, int64(v.ID))

			ch <- i
		}(v)
	}

	// closer
	go func() {
		wg.Wait()
		close(ch)
	}()

	videos := make([]*video2.Video, 0)
	for v := range ch {
		if v.err != nil {
			return nil, v.err
		}
		videos = append(videos, v.video)
	}

	return videos, nil
}
