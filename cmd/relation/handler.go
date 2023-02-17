package main

import (
	"context"
	"mini-tiktok-backend/cmd/relation/pack"
	"mini-tiktok-backend/cmd/relation/service"
	relation "mini-tiktok-backend/kitex_gen/relation"
	"mini-tiktok-backend/pkg/errno"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	resp = new(relation.RelationActionResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp,nil
	}

	err = service.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp,nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// GetFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *relation.GetFollowListRequest) (resp *relation.GetFollowListResponse, err error) {
	resp = new(relation.GetFollowListResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp,nil
	}

	users,err := service.NewGetFollowListService(ctx).GetFollowList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp,nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserList = users
	return
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *relation.GetFollowerListRequest) (resp *relation.GetFollowerListResponse, err error) {
	resp = new(relation.GetFollowerListResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp,nil
	}

	users,err := service.NewGetFollowerListService(ctx).GetFollowerList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp,nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserList = users
	return
}

// GetFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriendList(ctx context.Context, req *relation.GetFriendListRequest) (resp *relation.GetFriendListResponse, err error) {
	resp = new(relation.GetFriendListResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp,nil
	}

	users,err := service.NewFriendListService(ctx).GetFriendList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp,nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserList = users
	return
}

// MessageAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MessageAction(ctx context.Context, req *relation.MessageActionRequest) (resp *relation.MessageActionResponse, err error) {
	resp = new(relation.MessageActionResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp,nil
	}

	err = service.NewMessageActionService(ctx).MessageAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp,nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// MessageChat implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MessageChat(ctx context.Context, req *relation.MessageChatRequest) (resp *relation.MessageChatResponse, err error) {
	resp = new(relation.MessageChatResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp,nil
	}
	messages,err := service.NewMessageChatService(ctx).MessageChat(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp,nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.MessageList = messages
	return
}
