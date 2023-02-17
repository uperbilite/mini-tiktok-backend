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

func (s *MessageChatService) MessageChat(req *relation.MessageChatRequest) ([]*relation.Message,error) {
	userId,ToUserId := req.GetUserId(),req.GetToUserId()
	res,err := db.QueryMessageBothId(s.ctx,userId,ToUserId)
	if err != nil {
		return nil, err
	}
	message := make([]*relation.Message,len(res))
	for i,m := range res {
		message[i] = &relation.Message{
			Id: m.Id,
			CreateTime: m.CreatedAt.Unix(),
			Content: m.Content,
		}
	}
	return message,nil
}