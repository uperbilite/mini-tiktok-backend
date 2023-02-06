package pack

import (
	"mini-tiktok-backend/cmd/publish/dal/db"
	"mini-tiktok-backend/kitex_gen/publish"
)

func Video(v *db.Video) *publish.Video {
	if v == nil {
		return nil
	}
	return &publish.Video{
		Id:            int64(v.ID),
		Author:        nil, // TODO: Get author info from user service
		PlayUrl:       v.PlayURL,
		CoverUrl:      v.CoverURL,
		FavoriteCount: int64(v.FavouriteCount),
		CommentCount:  int64(v.CommentCount),
		IsFavorite:    false, // TODO: Get is favourite from favourite service
		Title:         v.Title,
	}
}
