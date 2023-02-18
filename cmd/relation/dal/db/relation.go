package db

import (
	"context"
	"gorm.io/gorm"
	"mini-tiktok-backend/pkg/consts"
)

type Follow struct {
	gorm.Model
	UserId int64 `json:"user_id"`
	FollowId int64 `json:"follow_id"`
}

type Follower struct {
	gorm.Model
	UserId int64 `json:"user_id"`
	FollowerId int64 `json:"follower_id"`
}

func (f *Follow) TableName() string {
	return consts.FollowTableName
}

func (f *Follower) TableName() string {
	return consts.FollowerTableName
}

// QueryFollowById query list of user info
func QueryFollowById(ctx context.Context, userId int64) ([]*Follow, error) {
	var res []*Follow
	if err := DB.WithContext(ctx).
		Where("user_id = ?", userId).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryFollowerById(ctx context.Context, userId int64) ([]*Follower, error) {
	var res []*Follower
	if err := DB.WithContext(ctx).
		Where("user_id = ?",userId).
		Find(&res).Error; err != nil {
			return nil, err
	}
	return res,nil
}

func QueryFriendById(ctx context.Context, userId int64) ([]*Follow, error) {
	var res []*Follow
	if err := DB.WithContext(ctx).
		Joins("JOIN followers ON follows.user_id = ? AND follows.follow_id = followers.follower_id",userId).
		Find(&res).Error; err != nil {
			return nil, err
	}
	return res,nil
}

func FollowUser(ctx context.Context, userId,targetUserId int64) error {
	err := DB.WithContext(ctx).Create(&Follow{UserId: userId,FollowId: targetUserId}).Error
	if err != nil {
		return err
	}
	err = DB.WithContext(ctx).Create(&Follower{UserId: targetUserId,FollowerId: userId}).Error
	if err != nil {
		return err
	}
	return nil
}

func CountFollow(ctx context.Context, userId int64) (follows int64,err error) {
	err = DB.WithContext(ctx).Model(&Follow{}).Where("user_id = ?",userId).Count(&follows).Error
	if err != nil {
		return 0, err
	}
	return
}

func CountFollower(ctx  context.Context,userId int64) (followers int64,err error) {
	err = DB.WithContext(ctx).Model(&Follower{}).Where("user_id = ?",userId).Count(&followers).Error
	if err != nil {
		return 0, err
	}
	return
}

func IsFollow(ctx context.Context, userId, toUserId int64) (bool, error) {
	err := DB.WithContext(ctx).
		Where("user_id = ? AND follow_id = ?",userId,toUserId).
		First(&Follow{}).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true,nil
}

