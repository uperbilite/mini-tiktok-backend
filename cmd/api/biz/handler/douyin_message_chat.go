package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	api_relation "mini-tiktok-backend/cmd/api/biz/model/api/relation"
	"mini-tiktok-backend/cmd/api/biz/mw"
	"mini-tiktok-backend/cmd/api/biz/rpc"
	"mini-tiktok-backend/kitex_gen/relation"
	pkg_consts "mini-tiktok-backend/pkg/consts"
	"mini-tiktok-backend/pkg/errno"
)

// DouyinMessageChat .
// @router /douyin/message/chat/ [GET]
func DouyinMessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api_relation.DouyinMessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	user, _ := c.Get(pkg_consts.IdentityKey)
	messageList, err := rpc.MessageChat(ctx, &relation.MessageChatRequest{
		UserId: user.(*mw.User).UserId,
		ToUserId: req.ToUserID,
	})
	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	SendResponse(c, errno.Success, utils.H{
		"message_list": messageList,
	})
}
