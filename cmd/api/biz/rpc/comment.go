package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"mini-tiktok-backend/kitex_gen/comment"
	"mini-tiktok-backend/kitex_gen/comment/commentservice"
	"mini-tiktok-backend/pkg/consts"
	"mini-tiktok-backend/pkg/errno"
)

var commentClient commentservice.Client

func initComment() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	c, err := commentservice.NewClient(
		consts.CommentServiceName,
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	commentClient = c
}

func CreateComment(ctx context.Context, req *comment.CreateCommentRequest) (*comment.Comment, error) {
	resp, err := commentClient.CreateComment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.Comment, nil
}

func DeleteComment(ctx context.Context, req *comment.DeleteCommentRequest) error {
	resp, err := commentClient.DeleteComment(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}
