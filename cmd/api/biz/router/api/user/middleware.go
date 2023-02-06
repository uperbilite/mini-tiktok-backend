// Code generated by hertz generator.

package User

import (
	"github.com/cloudwego/hertz/pkg/app"
	"mini-tiktok-backend/cmd/api/biz/mw"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		mw.AccessLog(),
	}
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinuserMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		// use jwt mw
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinuserloginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinuserregisterMw() []app.HandlerFunc {
	// your code...
	return nil
}
