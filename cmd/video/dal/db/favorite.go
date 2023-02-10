package db

import (
	"context"
	"strconv"
	"strings"
)

// GetFavoriteKey Key format is "favorite:{video_id}"
func GetFavoriteKey(videoId int64) string {
	var res strings.Builder
	res.WriteString("favorite:")
	res.WriteString(strconv.FormatInt(videoId, 10))
	return res.String()
}

func GetFavoriteCount(ctx context.Context, videoId int64) (int64, error) {
	res := RDB.Get(ctx, GetFavoriteKey(videoId))
	if res == nil {
		return 0, nil
	}
	return res.Int64()
}
