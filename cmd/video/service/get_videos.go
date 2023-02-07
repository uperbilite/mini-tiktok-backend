package service

import (
	"context"
	"mini-tiktok-backend/cmd/video/dal/db"
	"mini-tiktok-backend/cmd/video/pack"
	"mini-tiktok-backend/cmd/video/rpc"
	"mini-tiktok-backend/kitex_gen/favorite"
	"mini-tiktok-backend/kitex_gen/user"
	video2 "mini-tiktok-backend/kitex_gen/video"
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

	videos := make([]*video2.Video, 0)

	for _, v := range vs {
		video := pack.Video(v)

		resp, _ := rpc.QueryUser(s.ctx, &user.QueryUserRequest{
			UserId:       req.UserId,
			TargetUserId: v.AuthorId,
		})
		// TODO: err handle
		video.Author = pack.User(resp)

		// get is_favorite
		isFavorite, _ := rpc.GetIsFavorite(s.ctx, &favorite.GetIsFavoriteRequest{
			UserId:  req.UserId,
			VideoId: int64(v.ID),
		})
		// TODO: err handle
		video.IsFavorite = isFavorite

		// get favorite count
		favoriteCount, _ := rpc.GetFavoriteCount(s.ctx, &favorite.GetFavoriteCountRequest{
			VideoId: int64(v.ID),
		})
		// TODO: err handle
		video.FavoriteCount = favoriteCount

		videos = append(videos, video)
	}

	return videos, nil
}
