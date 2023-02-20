package service

import (
	"context"
	"mini-tiktok-backend/cmd/relation/dal/db"
	"mini-tiktok-backend/kitex_gen/relation"
)

type MessageActionService struct {
	ctx context.Context
}

func NewMessageActionService(ctx context.Context) *MessageActionService {
	return &MessageActionService{ctx}
}

func (s *MessageActionService) MessageAction(req *relation.MessageActionRequest) error {
	message := &db.Message{
		UserId: req.GetUserId(),
		ToUserId: req.GetToUserId(),
		Content: req.GetContent(),
	}
	actionType := req.GetActionType()
	if actionType == 1 {
		if err := db.CreateMessage(s.ctx,message); err != nil {
			return err
		}
	}
	return nil
}
