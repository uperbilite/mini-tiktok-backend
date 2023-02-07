package service

import (
	"context"
	"mini-tiktok-backend/cmd/publish/dal/db"
	"mini-tiktok-backend/cmd/publish/pack"
	"mini-tiktok-backend/cmd/publish/rpc"
	"mini-tiktok-backend/kitex_gen/favorite"
	"mini-tiktok-backend/kitex_gen/publish"
	"mini-tiktok-backend/kitex_gen/user"
)

type GetPublishListService struct {
	ctx context.Context
}

func NewGetPublishListService(ctx context.Context) *GetPublishListService {
	return &GetPublishListService{ctx: ctx}
}

func (s *GetPublishListService) GetPublishList(req *publish.GetPublishListRequest) ([]*publish.Video, error) {
	vs, err := db.GetVideosByAuthorId(s.ctx, req.TargetUserId)
	if err != nil {
		return nil, err
	}

	videos := make([]*publish.Video, 0)

	// TODO: get user info from video author id, using UserId and TargetUserId
	for _, v := range vs {
		video := pack.Video(v)

		resp, _ := rpc.QueryUser(s.ctx, &user.QueryUserRequest{
			UserId:       req.UserId,
			TargetUserId: req.TargetUserId,
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
