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

func TestGetFollowListService_GetFollowList(t *testing.T) {
	getFollowListService := NewGetFollowListService(context.Background())
	follow,err := getFollowListService.GetFollowList(&relation.GetFollowListRequest{UserId: 1,TargetUserId: 1})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(follow)
}
