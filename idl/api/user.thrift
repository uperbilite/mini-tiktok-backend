namespace go api.user

struct User {
    1: required i64 id
    2: required string name
    3: required i64 follow_count
    4: required i64 follower_count
    5: required bool is_follow
}

struct DouyinUserLoginRequest {
    1: required string username (api.query="username")
    2: required string password (api.query="password")
}

struct DouyinUserLoginResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required i64 user_id
    4: required string token
}

struct DouyinUserRegisterRequest {
    1: required string username (api.query="username", api.vd="len($) > 0 && len($) <= 32")
    2: required string password (api.query="password", api.vd="len($) > 0 && len($) <= 32")
}

struct DouyinUserRegisterResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required i64 user_id
    4: required string token
}

struct DouyinUserRequest {
    1: required i64 user_id (api.query="user_id")
    2: required string token (api.query="token")
}

struct DouyinUserResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required User user
}

service ApiUserService {
    DouyinUserLoginResponse DouyinUserLogin(1: DouyinUserLoginRequest req) (api.post="/douyin/user/login/")
    DouyinUserRegisterResponse DouyinUserRegister(1: DouyinUserRegisterRequest req) (api.post="/douyin/user/register/")
    DouyinUserResponse DouyinUser(1: DouyinUserRequest req) (api.get="/douyin/user/")
}