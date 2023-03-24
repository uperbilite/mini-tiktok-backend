// Code generated by hertz generator.

package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"mini-tiktok-backend/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	api_relation "mini-tiktok-backend/cmd/api/biz/model/api/relation"
)

// DouyinRelationAction .
// @router /douyin/relation/action/ [POST]
func DouyinRelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api_relation.DouyinRelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	/*user, _ := c.Get(pkg_consts.IdentityKey)

	err = rpc.RelationAction(ctx, &relation.RelationActionRequest{
		UserId:     user.(*mw.User).UserId,
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
	})

	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}*/

	SendResponse(c, errno.Success, utils.H{})
}
