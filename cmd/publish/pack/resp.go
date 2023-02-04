package pack

import (
	"errors"
	"mini-tiktok-backend/kitex_gen/publish"
	"mini-tiktok-backend/pkg/errno"
	"time"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *publish.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *publish.BaseResp {
	return &publish.BaseResp{
		StatusCode:  err.ErrCode,
		StatusMsg:   err.ErrMsg,
		ServiceTime: time.Now().Unix(),
	}
}
