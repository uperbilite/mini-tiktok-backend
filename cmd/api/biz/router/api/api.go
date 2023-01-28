// Code generated by hertz generator. DO NOT EDIT.

package Api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	api "mini-tiktok-backend/cmd/api/biz/handler/api"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_user := _douyin.Group("/user", _userMw()...)
			_user.GET("/", append(_queryuserMw(), api.QueryUser)...)
			{
				_login := _user.Group("/login", _loginMw()...)
				_login.POST("/", append(_checkuserMw(), api.CheckUser)...)
			}
			{
				_register := _user.Group("/register", _registerMw()...)
				_register.POST("/", append(_createuserMw(), api.CreateUser)...)
			}
		}
	}
}
