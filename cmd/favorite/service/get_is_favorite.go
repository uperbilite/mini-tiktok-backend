package service

import (
	"context"
	"mini-tiktok-backend/cmd/favorite/dal/db"
	"mini-tiktok-backend/kitex_gen/favorite"
)

type GetFavoriteService struct {
	ctx context.Context
}

func NewGetFavoriteService(ctx context.Context) *GetFavoriteService {
	return &GetFavoriteService{
		ctx: ctx,
	}
}

func (s *GetFavoriteService) GetIsFavorite(req *favorite.GetIsFavoriteRequest) (bool, error) {
	if req.UserId == 0 { // no login user
		return false, nil
	}

	favorites, err := db.QueryFavorite(s.ctx, req.UserId, req.VideoId)
	if err != nil {
		return false, err
	}
	if len(favorites) == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
