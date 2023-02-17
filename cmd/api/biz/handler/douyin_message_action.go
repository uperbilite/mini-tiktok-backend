package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	api_relation "mini-tiktok-backend/cmd/api/biz/model/api/relation"
	"mini-tiktok-backend/cmd/api/biz/mw"
	"mini-tiktok-backend/cmd/api/biz/rpc"
	"mini-tiktok-backend/kitex_gen/relation"
	pkg_consts "mini-tiktok-backend/pkg/consts"
	"mini-tiktok-backend/pkg/errno"
)

// DouyinMessageAction .
// @router /douyin/message/action/ [GET]
func DouyinMessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api_relation.DouyinMessageActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	user, _ := c.Get(pkg_consts.IdentityKey)

	err = rpc.MessageAction(ctx, &relation.MessageActionRequest{
		UserId:     user.(*mw.User).UserId,
		ToUserId:   req.ToUserID,
		Content: req.Content,
	})

	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	SendResponse(c, errno.Success, utils.H{})
}