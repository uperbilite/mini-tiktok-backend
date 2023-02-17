package main

import (
	"context"
	relation "mini-tiktok-backend/kitex_gen/relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *relation.GetFollowListRequest) (resp *relation.GetFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *relation.GetFollowerListRequest) (resp *relation.GetFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriendList(ctx context.Context, req *relation.GetFriendListRequest) (resp *relation.GetFriendListResponse, err error) {
	// TODO: Your code here...
	return
}

// MessageAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MessageAction(ctx context.Context, req *relation.MessageActionRequest) (resp *relation.MessageActionResponse, err error) {
	// TODO: Your code here...
	return
}

// MessageChat implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MessageChat(ctx context.Context, req *relation.MessageChatRequest) (resp *relation.MessageChatResponse, err error) {
	// TODO: Your code here...
	return
}
