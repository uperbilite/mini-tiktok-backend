package service

import (
	"context"
	"mini-tiktok-backend/cmd/relation/dal/db"
	"mini-tiktok-backend/kitex_gen/relation"
	"testing"
)

func init() {
	db.Init()
}

func TestMessageActionService_MessageAction(t *testing.T) {
	messageActionService := NewMessageActionService(context.Background())
	err := messageActionService.MessageAction(&relation.MessageActionRequest{UserId: 1,ToUserId: 2,Content: "11111"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
