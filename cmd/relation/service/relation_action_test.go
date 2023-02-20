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

func TestRelationActionService_RelationAction1(t *testing.T) {
	ctx := context.Background()
	relationActionService := NewRelationActionService(ctx)
	err := relationActionService.RelationAction(&relation.RelationActionRequest{UserId: 2, ToUserId: 1, ActionType: 1})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
func TestRelationActionService_RelationAction2(t *testing.T) {
	ctx := context.Background()
	relationActionService := NewRelationActionService(ctx)
	err := relationActionService.RelationAction(&relation.RelationActionRequest{UserId: 1, ToUserId: 2, ActionType: 2})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}
