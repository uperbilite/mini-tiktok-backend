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

func TestIsFollowToUserService_IsFollowToUser(t *testing.T) {
	isFollowToUserService := NewIsFollowToUserService(context.Background())
	is_follow,err := isFollowToUserService.IsFollowToUser(&relation.IsFollowToUserRequest{UserId: 1,ToUserId: 2})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("user1 is follow to user2: ",is_follow)
}
