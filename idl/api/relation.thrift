namespace go api.relation

struct User {
    1: required i64 id
    2: required string name
    3: required i64 follow_count
    4: required i64 follower_count
    5: required bool is_follow
}

struct Message {
    1: required i64 id
    2: required i64 to_user_id
    3: required i64 from_user_id
    4: required string content
    5: required i64 create_time
}

struct DouyinRelationActionRequest {
    1: required string token (api.query="token")
    2: required i64 to_user_id (api.query="to_user_id")
    3: required i32 action_type (api.query="action_type" api.vd="in($, 1, 2)")
}

struct DouyinRelationActionResponse {
    1: required i32 status_code
    2: required string status_msg
}

struct DouyinRelationFollowListRequest {
    1: required i64 user_id (api.query="user_id")
    2: required string token (api.query="token")
}

struct DouyinRelationFollowListResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required list<User> user_list
}

struct DouyinRelationFollowerListRequest {
    1: required i64 user_id (api.query="user_id")
    2: required string token (api.query="token")
}

struct DouyinRelationFollowerListResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required list<User> user_list
}

struct DouyinRelationFriendListRequest {
    1: required i64 user_id (api.query="user_id")
    2: required string token (api.query="token")
}

struct DouyinRelationFriendListResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required list<User> user_list
}

struct DouyinMessageActionRequest {
    1: required string token (api.query="token")
    2: required i64 to_user_id (api.query="to_user_id")
    3: required i32 action_type (api.query="action_type" api.vd="in($, 1)")
    4: required string content (api.query="content")
}

struct DouyinMessageActionResponse {
    1: required i32 status_code
    2: required string status_msg
}

struct DouyinMessageChatRequest {
    1: required string token (api.query="token")
    2: required i64 to_user_id (api.query="to_user_id")
}

struct DouyinMessageChatResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required list<Message> message_list
}

service ApiRelationService {
    DouyinRelationActionResponse DouyinRelationAction(1: DouyinRelationActionRequest req) (api.post="/douyin/relation/action/")
    DouyinRelationFollowListResponse DouyinRelationFollowList(1: DouyinRelationFollowListRequest req) (api.get="/douyin/relation/follow/list/")
    DouyinRelationFollowerListResponse DouyinRelationFollowerList(1: DouyinRelationFollowerListRequest req) (api.get="/douyin/relation/follower/list/")
    DouyinRelationFriendListResponse DouyinRelationFriendList(1: DouyinRelationFriendListRequest req) (api.get="/douyin/relation/friend/list/")
    DouyinMessageActionResponse DouyinMessageAction(1: DouyinMessageActionRequest req) (api.post="/douyin/message/action/")
    DouyinMessageChatResponse DouyinMessageChat(1: DouyinMessageChatRequest req) (api.get="/douyin/message/chat/")
}