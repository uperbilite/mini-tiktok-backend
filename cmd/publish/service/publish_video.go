package service

import (
	"bytes"
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gofrs/uuid"
	"mini-tiktok-backend/cmd/publish/dal/db"
	"mini-tiktok-backend/kitex_gen/publish"
	"mini-tiktok-backend/pkg/consts"
	"strings"
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

	videoUid, _ := uuid.NewV4()
	videoUidString := videoUid.String()

	var videoObjectKey strings.Builder
	videoObjectKey.Grow(consts.ObjKeyLen)
	videoObjectKey.WriteString("video/")
	videoObjectKey.WriteString(videoUidString)
	videoObjectKey.WriteString(".mp4")

	err = bucket.PutObject(videoObjectKey.String(), bytes.NewReader(req.Data))
	if err != nil {
		return err
	}

	// TODO: gen signed url from video url, remove cover url in favorites table.
	return db.CreateVideo(s.ctx, &db.Video{
		AuthorId: req.UserId,
		PlayURL:  videoObjectKey.String(),
		CoverURL: "",
		Title:    req.Title,
	})

}
