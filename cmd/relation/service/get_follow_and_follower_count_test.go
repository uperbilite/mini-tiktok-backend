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

func TestGetFollowAndFollowerCountService_GetFollowAndFollowerCount(t *testing.T) {
	getFollowAndFollowerCountService := NewGetFollowAndFollowerCountService(context.Background())
	follows, followers, err := getFollowAndFollowerCountService.GetFollowAndFollowerCount(&relation.GetFollowAndFollowerCountRequest{UserId: 1})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("follows = %d,followers = %d", follows, followers)
}
