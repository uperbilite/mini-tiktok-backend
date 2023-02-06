package errno

import (
	"errors"
	"fmt"
	"mini-tiktok-backend/kitex_gen/user"
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success    = NewErrNo(0, "Success")
	ServiceErr = NewErrNo(5000, "Service is unable to start successfully")
	ParamErr   = NewErrNo(5001, "Wrong parameters has been given")

	UserNotExistErr            = NewErrNo(int32(user.ErrCode_UserNotExistErrCode), "User not exist")
	UserAlreadyExistErr        = NewErrNo(int32(user.ErrCode_UserAlreadyExistErrCode), "User already exists")
	UserAuthorizationFailedErr = NewErrNo(int32(user.ErrCode_AuthorizationFailedErrCode), "User authorization failed")

	// TODO: add more error code, including publish error and user error
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
