package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"mini-tiktok-backend/pkg/consts"
	"strconv"
)

type Follow struct {
	gorm.Model
	FromId int64 `json:"from_id"`
	ToId   int64 `json:"to_id"`
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
		Where("to_id = ?", userId).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryFriendById(ctx context.Context, userId int64) ([]*Follow, error) {
	// TODO: optimize by raw sql
	var res []*Follow
	if err := DB.WithContext(ctx).
		Where("from_id IN (?) AND to_id = ?", DB.Model(&Follow{}).Select("to_id").Where("from_id = ?", userId), userId).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func FollowUser(ctx context.Context, fromId, toId int64) (err error) {
	err = DB.WithContext(ctx).Where("from_id = ? AND to_id = ?", fromId, toId).First(&Follow{}).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	if err == nil {
		return
	}
	err = DB.WithContext(ctx).Create(&Follow{FromId: fromId, ToId: toId}).Error
	if err != nil {
		return
	}
	return
}

func CancelFollow(ctx context.Context, fromId, toId int64) (err error) {
	err = DB.WithContext(ctx).Where("from_id = ? AND to_id = ?", fromId, toId).Delete(&Follow{}).Error
	if err != nil {
		return
	}
	return
}

func CountFollow(ctx context.Context, userId int64) (follows int64, err error) {
	err = DB.WithContext(ctx).Model(&Follow{}).Where("from_id = ?", userId).Count(&follows).Error
	if err != nil {
		return 0, err
	}
	return
}

func CountFollower(ctx context.Context, userId int64) (followers int64, err error) {
	err = DB.WithContext(ctx).Model(&Follow{}).Where("to_id = ?", userId).Count(&followers).Error
	if err != nil {
		return 0, err
	}
	return
}

func RemoveKeyFromRedis(ctx context.Context, fromId, toId int64) (err error) {
	err = RDB.Del(ctx, strconv.FormatInt(fromId, 10), strconv.FormatInt(toId, 10)).Err()
	return
}

func SetFollowNumAndFollowerNumToRedis(ctx context.Context, userId, follows, followers int64) error {
	err := RDB.HMSet(ctx, strconv.FormatInt(userId, 10), "follows", follows, "followers", followers).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetFollowNumAndFollowerNumFromRedis(ctx context.Context, userId int64) (int64, int64, error) {
	ff, err := RDB.HMGet(ctx, strconv.FormatInt(userId, 10), "follows", "followers").Result()
	if err != nil {
		return 0, 0, err
	}
	if ff[0] == nil || ff[1] == nil {
		return 0, 0, errors.New("record not found")
	}
	follows, _ := strconv.ParseInt(ff[0].(string), 10, 0)
	followers, _ := strconv.ParseInt(ff[1].(string), 10, 0)
	return follows, followers, nil
}

func IsFollow(ctx context.Context, userId, toUserId int64) (bool, error) {
	err := DB.WithContext(ctx).
		Where("from_id = ? AND to_id = ?", userId, toUserId).
		First(&Follow{}).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
