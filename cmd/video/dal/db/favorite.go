package db

import (
	"context"
	"strconv"
	"strings"
)

func GetVideoKey(videoId int64) string {
	var res strings.Builder
	res.WriteString("video:")
	res.WriteString(strconv.FormatInt(videoId, 10))
	return res.String()
}

func GetFavoriteCount(ctx context.Context, videoId int64) (int64, error) {
	res := RDB.Get(ctx, GetVideoKey(videoId))
	if res == nil {
		return 0, nil
	}
	return res.Int64()
}
