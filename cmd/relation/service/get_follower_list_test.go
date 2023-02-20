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

func TestGetFollowerListService_GetFollowerList(t *testing.T) {
	getFollowerListService := NewGetFollowerListService(context.Background())
	follower, err := getFollowerListService.GetFollowerList(&relation.GetFollowerListRequest{UserId: 1, TargetUserId: 1})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(follower)
}
