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
	err := DB.WithContext(ctx).Create(Follow{UserId: userId,FollowId: targetUserId}).Error
	if err != nil {
		return err
	}
	err = DB.WithContext(ctx).Create(&Follower{UserId: targetUserId,FollowerId: userId}).Error
	if err != nil {
		return err
	}
	return nil
}

