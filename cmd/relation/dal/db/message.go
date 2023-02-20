package db

import (
	"context"
	"gorm.io/gorm"
	"mini-tiktok-backend/pkg/consts"
)

type Message struct {
	gorm.Model
	Id int64 `json:"id"`
	UserId int64 `json:"user_id"`
	ToUserId int64 `json:"to_user_id"`
	Content string `json:"content"`
}

func (m *Message) TableName() string {
	return consts.MessageTableName
}

func CreateMessage(ctx context.Context, message *Message) error {
	return DB.WithContext(ctx).Create(message).Error
}

func QueryMessageBothId(ctx context.Context,userId,toUserId int64,preMsgTime int64) ([]*Message,error) {
	var me2You,you2Me []*Message
	if err := DB.WithContext(ctx).
		Where("user_id = ? AND to_user_id = ? AND create_at > ?",userId,toUserId,preMsgTime).
		Find(&me2You).Error; err  != nil {
		return nil, err
	}
	if err := DB.WithContext(ctx).
		Where("user_id = ? AND to_user_id = ? AND create_at > ?",toUserId,userId,preMsgTime).
		Find(&you2Me).Error; err != nil {
		return nil, err
	}
	res := make([]*Message,0,len(me2You) + len(you2Me))
	res = append(res,me2You...)
	res = append(res,you2Me...)
	return res,nil
}
