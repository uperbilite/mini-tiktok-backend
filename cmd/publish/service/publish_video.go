package service

import (
	"bytes"
	"context"
	"mini-tiktok-backend/cmd/publish/dal/db"
	"mini-tiktok-backend/kitex_gen/publish"
	"mini-tiktok-backend/pkg/consts"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type PublishVideoService struct {
	ctx context.Context
}

func NewPublishVideoService(ctx context.Context) *PublishVideoService {
	return &PublishVideoService{ctx: ctx}
}

func (s *PublishVideoService) PublishVideo(req *publish.PublishVideoRequest) error {
	// push video to OSS and get url
	client, err := oss.New(consts.OSSEndPoint, consts.AccessKeyId, consts.AccessKeySecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(consts.OSSBucketName)
	if err != nil {
		return err
	}

	err = bucket.PutObject(req.Title+".mp4", bytes.NewReader(req.Data)) // TODO: set key access
	if err != nil {
		return err
	}

	return db.CreateVideo(s.ctx, &db.Video{
		AuthorId: req.UserId,
		PlayURL:  consts.OSSResourceURL + req.Title + ".mp4",
		CoverURL: consts.OSSResourceURL + req.Title + ".jpeg", // TODO: get cover url
		Title:    req.Title,
	})

}
