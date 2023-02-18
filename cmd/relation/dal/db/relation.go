package db

import (
	"context"
	"gorm.io/gorm"
	"mini-tiktok-backend/pkg/consts"
)

type Follow struct {
	gorm.Model
	FromId int64 `json:"from_id"`
	ToId int64 `json:"to_id"`
}

func (f *Follow) TableName() string {
	return consts.FollowTableName
}

// QueryFollowById query list of user info
func QueryFollowById(ctx context.Context, userId int64) ([]*Follow, error) {
	var res []*Follow
	if err := DB.WithContext(ctx).
		Where("from_id = ?", userId).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryFollowerById(ctx context.Context, userId int64) ([]*Follow, error) {
	var res []*Follow
	if err := DB.WithContext(ctx).
		Where("to_id = ?",userId).
		Find(&res).Error; err != nil {
			return nil, err
	}
	return res,nil
}

func QueryFriendById(ctx context.Context, userId int64) ([]*Follow, error) {
	var res []*Follow
	if err := DB.WithContext(ctx).
		Where("from_id IN ? AND to_id = ?",DB.Model(&Follow{}).Select("to_id").Where("from_id = ?",userId),userId).
		Find(&res).Error; err != nil {
			return nil, err
	}
	return res,nil
}

func FollowUser(ctx context.Context, fromId,toId int64) (err error) {
	err = DB.WithContext(ctx).Create(&Follow{FromId: fromId,ToId: toId}).Error
	if err != nil {
		return
	}
	return
}

func CancelFollow(ctx context.Context, fromId, toId int64) (err error) {
	err = DB.WithContext(ctx).Delete(&Follow{FromId: fromId,ToId: toId}).Error
	if err != nil {
		return
	}
	return
}

func CountFollow(ctx context.Context, userId int64) (follows int64,err error) {
	err = DB.WithContext(ctx).Model(&Follow{}).Where("from_id = ?",userId).Count(&follows).Error
	if err != nil {
		return 0, err
	}
	return
}

func CountFollower(ctx  context.Context,userId int64) (followers int64,err error) {
	err = DB.WithContext(ctx).Model(&Follow{}).Where("to_id = ?",userId).Count(&followers).Error
	if err != nil {
		return 0, err
	}
	return
}

func IsFollow(ctx context.Context, userId, toUserId int64) (bool, error) {
	err := DB.WithContext(ctx).
		Where("from_id = ? AND to_id = ?",userId,toUserId).
		First(&Follow{}).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true,nil
}

