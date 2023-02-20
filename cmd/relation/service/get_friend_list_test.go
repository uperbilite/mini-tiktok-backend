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

func TestGetFriendListService_GetFriendList(t *testing.T) {
	getFriendListService := NewFriendListService(context.Background())
	friend, err := getFriendListService.GetFriendList(&relation.GetFriendListRequest{UserId: 1, TargetUserId: 1})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(friend)
}
