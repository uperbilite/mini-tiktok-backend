// Code generated by hertz generator. DO NOT EDIT.

package Relation

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	handler "mini-tiktok-backend/cmd/api/biz/handler"
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
			_relation := _douyin.Group("/relation", _relationMw()...)
			{
				_action := _relation.Group("/action", _actionMw()...)
				_action.POST("/", append(_douyinrelation_ctionMw(), handler.DouyinRelationAction)...)
			}
			{
				_follow := _relation.Group("/follow", _followMw()...)
				{
					_list := _follow.Group("/list", _listMw()...)
					_list.GET("/", append(_douyinrelationfollowlistMw(), handler.DouyinRelationFollowList)...)
				}
			}
			{
				_follower := _relation.Group("/follower", _followerMw()...)
				{
					_list0 := _follower.Group("/list", _list0Mw()...)
					_list0.GET("/", append(_douyinrelationfollowerlistMw(), handler.DouyinRelationFollowerList)...)
				}
			}
			{
				_friend := _relation.Group("/friend", _friendMw()...)
				{
					_list1 := _friend.Group("/list", _list1Mw()...)
					_list1.GET("/", append(_douyinrelationfriendlistMw(), handler.DouyinRelationFriendList)...)
				}
			}
		}
	}
}
