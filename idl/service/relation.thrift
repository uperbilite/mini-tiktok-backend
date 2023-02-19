namespace go relation

struct BaseResp {
    1: required i32 status_code
    2: required string status_msg
    3: required i64 service_time
}

struct User {
    1: required i64 id
    2: required string name
    3: required i64 follow_count
    4: required i64 follower_count
    5: required bool is_follow
}

struct Message {
    1: required i64 id
    2: required string content
    3: required string create_time
}

struct RelationActionRequest {
    1: required i64 user_id (vt.gt = "0")
    2: required i64 to_user_id (vt.gt = "0")
    3: required i32 action_type (vt.in = "1", vt.in = "2")
}

struct RelationActionResponse {
    1: required BaseResp base_resp
}

struct GetFollowListRequest {
    1: required i64 user_id (vt.gt = "0") // 0表示用户未登录
    2: required i64 target_user_id (vt.gt = "0")
}

struct GetFollowListResponse {
    1: required BaseResp base_resp
    2: required list<User> user_list
}

struct GetFollowerListRequest {
    1: required i64 user_id (vt.gt = "0") // 0表示用户未登录
    2: required i64 target_user_id (vt.gt = "0")
}

struct GetFollowerListResponse {
    1: required BaseResp base_resp
    2: required list<User> user_list
}

struct GetFriendListRequest {
    1: required i64 user_id (vt.gt = "0") // 0表示用户未登录
    2: required i64 target_user_id (vt.gt = "0")
}

struct GetFriendListResponse {
    1: required BaseResp base_resp
    2: required list<User> user_list
}

struct GetFollowAndFollowerCountRequest {
    1: required i64 user_id
}

struct GetFollowAndFollowerCountResponse {
    1: required BaseResp base_resp
    2: required i64 follows
    3: required i64 followers
}

struct IsFollowToUserRequest {
    1: required i64 user_id
    2: required i64 to_user_id
}

struct IsFollowToUserResponse {
    1: required BaseResp base_resp
    2: required bool is_follow
}

struct MessageActionRequest {
    1: required i64 user_id (vt.gt = "0")
    2: required i64 to_user_id (vt.gt = "0")
    3: required string content (vt.min_size = "1")
}

struct MessageActionResponse {
    1: required BaseResp base_resp
}

struct MessageChatRequest {
    1: required i64 user_id (vt.gt = "0")
    2: required i64 to_user_id (vt.gt = "0")
}

struct MessageChatResponse {
    1: required BaseResp base_resp
    2: required list<Message> message_list
}

service RelationService {
    RelationActionResponse RelationAction(1: RelationActionRequest req)
    GetFollowListResponse GetFollowList(1: GetFollowListRequest req)
    GetFollowerListResponse GetFollowerList(1: GetFollowerListRequest req)
    GetFriendListResponse GetFriendList(1: GetFriendListRequest req)
    MessageActionResponse MessageAction(1: MessageActionRequest req)
    MessageChatResponse MessageChat(1: MessageChatRequest req)
    GetFollowAndFollowerCountResponse GetFollowAndFollowerCount(1: GetFollowAndFollowerCountRequest req)
    IsFollowToUserResponse IsFollowToUser(1: IsFollowToUserRequest req)
}







