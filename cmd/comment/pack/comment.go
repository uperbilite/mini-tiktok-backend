package pack

import (
	"mini-tiktok-backend/cmd/comment/dal/db"
	"mini-tiktok-backend/kitex_gen/comment"
)

func Comment(c *db.Comment) *comment.Comment {
	if c == nil {
		return nil
	}
	return &comment.Comment{
		Id:         int64(c.ID),
		User:       nil,
		Content:    c.Content,
		CreateDate: c.CreateDate,
	}
}
