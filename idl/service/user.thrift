namespace go user

enum ErrCode {
    UserNotExistErrCode        = 10001
    UserAlreadyExistErrCode    = 10002
    AuthorizationFailedErrCode = 10003
}

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

struct CheckUserRequest {
    1: required string username (vt.min_size = "1")
    2: required string password (vt.min_size = "1")
}

struct CheckUserResponse {
    1: required BaseResp base_resp
    2: required i64 user_id
}

struct CreateUserRequest {
    1: required string username (vt.min_size = "1")
    2: required string password (vt.min_size = "1")
}

struct CreateUserResponse {
    1: required BaseResp base_resp
}

struct QueryUserRequest {
    1: required i64 user_id // 0表示用户未登录
    2: required i64 target_user_id (vt.gt = "0")
}

struct QueryUserResponse {
    1: required BaseResp base_resp
    2: required User user // user_id为0不获取is_followed
}

service UserService {
    CheckUserResponse CheckUser(1: CheckUserRequest req)
    CreateUserResponse CreateUser(1: CreateUserRequest req)
    QueryUserResponse QueryUser(1: QueryUserRequest req)
}