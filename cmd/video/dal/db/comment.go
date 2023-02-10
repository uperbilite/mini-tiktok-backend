package db

import (
	"context"
	"strconv"
	"strings"
)

// GetCommentKey Key format is "comment:{video_id}"
func GetCommentKey(videoId int64) string {
	var res strings.Builder
	res.WriteString("comment:")
	res.WriteString(strconv.FormatInt(videoId, 10))
	return res.String()
}

func GetCommentCount(ctx context.Context, videoId int64) (int64, error) {
	res := RDB.Get(ctx, GetCommentKey(videoId))
	if res == nil {
		return 0, nil
	}
	return res.Int64()
}
