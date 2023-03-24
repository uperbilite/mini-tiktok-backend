package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type DistributedLock struct {
	TTL             int
	RandomValue     uint64
	Key             string
	TryLockInterval time.Duration
	watchDog        chan bool
}

func (l *DistributedLock) TryLock(ctx context.Context) error {
	cmd := RDB.Do(ctx, "set", l.Key, l.RandomValue, "px", l.TTL, "nx")
	if cmd.Err() != nil {
		return cmd.Err()
	}
	go l.startWatchDog(ctx)
	return nil
}

func (l *DistributedLock) Unlock(ctx context.Context) error {
	script, err := os.ReadFile("unlock.lua")
	if err != nil {
		return err
	}
	lua := redis.NewScript(string(script))

	_, err = lua.Run(ctx, RDB, []string{l.Key}, l.RandomValue).Result()
	close(l.watchDog)
	return err
}

func (l *DistributedLock) Lock(ctx context.Context) error {
	// 尝试加锁
	err := l.TryLock(ctx)
	if err == nil {
		return nil
	}
	// 加锁失败，不断尝试
	ticker := time.NewTicker(l.TryLockInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			// 超时
			return nil
		case <-ticker.C:
			// 重新尝试加锁
			err := l.TryLock(ctx)
			if err == nil {
				return nil
			}
		}
	}
}

func (l *DistributedLock) startWatchDog(ctx context.Context) {
	t := l.TTL * 500
	deleteTime := time.Duration(t)
	ticker := time.NewTicker(deleteTime)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			// 延长锁的过期时间
			ok, err := RDB.Do(ctx, "expire", l.Key, l.TTL).Result()
			// 异常或锁已经不存在则不再续期
			if (err != nil) || ok == 0 {
				return
			}
		case <-l.watchDog:
			// 已经解锁
			return
		}
	}
}

func NewFavoriteKeyLock(userId int64, prefix string) DistributedLock {
	value := rand.Intn(time.Now().Nanosecond())
	deleteTime := time.Duration(100)
	watchDog := make(chan bool)
	return DistributedLock{
		Key:             prefix + strconv.FormatInt(userId, 10),
		RandomValue:     uint64(value),
		TTL:             500,
		TryLockInterval: deleteTime,
		watchDog:        watchDog,
	}
}
