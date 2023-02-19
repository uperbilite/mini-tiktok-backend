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

func TestMessageChatService_MessageChat(t *testing.T) {
	messageChatService := NewMessageChatService(context.Background())
	messages,err := messageChatService.MessageChat(&relation.MessageChatRequest{UserId: 1,ToUserId: 2})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(messages)
}
