package pack

import (
	"errors"
	"mini-tiktok-backend/kitex_gen/favorite"
	"mini-tiktok-backend/pkg/errno"
	"time"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *favorite.BaseResp {
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

func baseResp(err errno.ErrNo) *favorite.BaseResp {
	return &favorite.BaseResp{
		StatusCode:  err.ErrCode,
		StatusMsg:   err.ErrMsg,
		ServiceTime: time.Now().Unix(),
	}
}

