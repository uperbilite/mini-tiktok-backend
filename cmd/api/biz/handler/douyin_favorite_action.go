// Code generated by hertz generator.

package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	api_favorite "mini-tiktok-backend/cmd/api/biz/model/api/favorite"
	"mini-tiktok-backend/cmd/api/biz/mw"
	"mini-tiktok-backend/cmd/api/biz/rpc"
	"mini-tiktok-backend/kitex_gen/favorite"
	pkg_consts "mini-tiktok-backend/pkg/consts"
	"mini-tiktok-backend/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
)

// DouyinFavoriteAction .
// @router /douyin/favorite/action/ [POST]
func DouyinFavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api_favorite.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	user, _ := c.Get(pkg_consts.IdentityKey)

	err = rpc.FavoriteAction(ctx, &favorite.FavoriteActionRequest{
		VideoId:    req.VideoID,
		UserId:     user.(*mw.User).UserId,
		ActionType: req.ActionType,
	})

	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	SendResponse(c, errno.Success, utils.H{})
}
