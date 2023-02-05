package service

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"mini-tiktok-backend/cmd/publish/dal/db"
	"mini-tiktok-backend/kitex_gen/publish"
	"mini-tiktok-backend/pkg/consts"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gofrs/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
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
	videoUrl := consts.OSSResourceURL + "video/" + videoUid.String() + ".mp4"
	err = bucket.PutObject("video/"+videoUid.String()+".mp4", bytes.NewReader(req.Data)) // TODO: set key access
	if err != nil {
		return err
	}
	cover, _ := GetCoverFromVideo(videoUrl)
	coverUrl := consts.OSSResourceURL + "cover/" + videoUid.String() + ".jpg"
	err = bucket.PutObject("cover/"+videoUid.String()+".jpg", bytes.NewReader(cover)) // TODO: set key access
	if err != nil {
		return err
	}

	return db.CreateVideo(s.ctx, &db.Video{
		AuthorId: req.UserId,
		PlayURL:  videoUrl,
		CoverURL: coverUrl,
		Title:    req.Title,
	})

}

func GetCoverFromVideo(videoUrl string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoUrl).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), err
}
