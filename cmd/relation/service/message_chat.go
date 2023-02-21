package service

import (
	"context"
	"mini-tiktok-backend/cmd/relation/dal/db"
	"mini-tiktok-backend/kitex_gen/relation"
)

type MessageChatService struct {
	ctx context.Context
}

func NewMessageChatService(ctx context.Context) *MessageChatService {
	return &MessageChatService{ctx}
}

func (s *MessageChatService) MessageChat(req *relation.MessageChatRequest) ([]*relation.Message, error) {
	userId, ToUserId := req.GetUserId(), req.GetToUserId()
	preMsgTime := req.GetPreMsgTime()
	res, err := db.QueryMessageBothId(s.ctx, userId, ToUserId, preMsgTime)
	if err != nil {
		return nil, err
	}
	message := make([]*relation.Message, len(res))
	for i, m := range res {
		message[i] = &relation.Message{
			Id:         m.Id,
			FromUserId: m.UserId,
			ToUserId:   m.ToUserId,
			Content:    m.Content,
			CreateTime: m.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}
	return message, nil
}
