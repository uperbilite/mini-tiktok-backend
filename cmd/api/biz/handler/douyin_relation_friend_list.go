// Code generated by hertz generator.

package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"mini-tiktok-backend/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	api_relation "mini-tiktok-backend/cmd/api/biz/model/api/relation"
)

// DouyinRelationFriendList .
// @router /douyin/relation/friend/list/ [GET]
func DouyinRelationFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api_relation.DouyinRelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	/*user, _ := c.Get(pkg_consts.IdentityKey)

	userList, err := rpc.GetFriendList(ctx, &relation.GetFriendListRequest{
		UserId:       user.(*mw.User).UserId,
		TargetUserId: req.UserID,
	})
	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}*/

	// var userList []*relation.User

	SendResponse(c, errno.Success, utils.H{})
}
