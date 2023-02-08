package main

import (
	"context"
	"mini-tiktok-backend/cmd/favorite/pack"
	"mini-tiktok-backend/cmd/favorite/service"
	favorite "mini-tiktok-backend/kitex_gen/favorite"
	"mini-tiktok-backend/pkg/errno"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	resp = new(favorite.FavoriteActionResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err) // pack err message
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)

	return resp, nil
}

func (s *FavoriteServiceImpl) GetIsFavorite(ctx context.Context, req *favorite.GetIsFavoriteRequest) (resp *favorite.GetIsFavoriteResponse, err error) {
	resp = new(favorite.GetIsFavoriteResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	isFavorite, err := service.NewGetIsFavoriteService(ctx).GetIsFavorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.IsFavorite = isFavorite

	return resp, nil
}

func (s *FavoriteServiceImpl) GetFavoriteCount(ctx context.Context, req *favorite.GetFavoriteCountRequest) (resp *favorite.GetFavoriteCountResponse, err error) {
	resp = new(favorite.GetFavoriteCountResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	favoriteCount, err := service.NewGetFavoriteCountService(ctx).GetFavoriteCount(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.FavoriteCount = favoriteCount

	return resp, nil
}

func (s *FavoriteServiceImpl) GetFavoriteList(ctx context.Context, req *favorite.GetFavoriteListRequest) (resp *favorite.GetFavoriteListResponse, err error) {
	resp = new(favorite.GetFavoriteListResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	videoList, err := service.NewGetFavoriteListService(ctx).GetFavoriteList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videoList

	return resp, nil
}
