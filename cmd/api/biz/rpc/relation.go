package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"mini-tiktok-backend/kitex_gen/relation"
	"mini-tiktok-backend/kitex_gen/relation/relationservice"
	"mini-tiktok-backend/pkg/consts"
	"mini-tiktok-backend/pkg/errno"
	"mini-tiktok-backend/pkg/mw"
)

var relationClient relationservice.Client

func initRelation() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := relationservice.NewClient(
		consts.FavoriteServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	relationClient = c
}

func RelationAction(ctx context.Context, req *relation.RelationActionRequest) error {
	resp, err := relationClient.RelationAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

func GetFollowList(ctx context.Context, req *relation.GetFollowListRequest) ([]*relation.User, error) {
	resp, err := relationClient.GetFollowList(ctx, req)
	if err != nil {
		return nil,err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil,errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserList,nil
}

func GetFollowerList(ctx context.Context, req *relation.GetFollowerListRequest) ([]*relation.User, error) {
	resp, err := relationClient.GetFollowerList(ctx, req)
	if err != nil {
		return nil,err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil,errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserList,nil
}

func GetFriendList(ctx context.Context, req *relation.GetFriendListRequest) ([]*relation.User, error) {
	resp, err := relationClient.GetFriendList(ctx, req)
	if err != nil {
		return nil,err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil,errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.UserList,nil
}

func MessageAction(ctx context.Context,req *relation.MessageActionRequest) error {
	resp,err := relationClient.MessageAction(ctx,req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

func MessageChat(ctx context.Context, req *relation.MessageChatRequest) ([]*relation.Message, error) {
	resp, err := relationClient.MessageChat(ctx,req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.MessageList,nil
}
