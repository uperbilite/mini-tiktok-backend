package handler

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"mini-tiktok-backend/pkg/errno"
)

// SendResponse Package response body with status_code and status_body.
func SendResponse(c *app.RequestContext, err error, data map[string]interface{}) {
	Err := errno.ConvertErr(err)
	data["status_code"] = int32(Err.ErrCode)
	data["status_msg"] = Err.ErrMsg
	c.JSON(consts.StatusOK, data)
}
