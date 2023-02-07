package handler

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/golang-jwt/jwt/v4"
	"mini-tiktok-backend/cmd/api/biz/mw"
	"mini-tiktok-backend/pkg/errno"
)

// SendResponse Package response body with status_code and status_body.
func SendResponse(c *app.RequestContext, err error, data map[string]interface{}) {
	Err := errno.ConvertErr(err)
	data["status_code"] = Err.ErrCode
	data["status_msg"] = Err.ErrMsg
	c.JSON(consts.StatusOK, data)
}

func GetClaimsFromTokenString(token string) (map[string]interface{}, error) {
	t, _ := mw.JwtMiddleware.ParseTokenString(token)
	claims := jwt.MapClaims{}
	for key, value := range t.Claims.(jwt.MapClaims) {
		claims[key] = value
	}
	return claims, nil
}
