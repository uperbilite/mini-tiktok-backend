namespace go api

struct User {
    1: required i64 id
    2: required string name
    // TODO: add follow and follow count, is follow status
}

struct DouyinUserLoginRequest {
    1: required string username (api.query="username")
    2: required string password (api.query="password")
}

struct DouyinUserLoginResponse {
    1: required i64 status_code
    2: required string status_msg
    3: required i64 user_id
    4: required string token
}

struct CreateUserRequest {
    1: required string username (api.query="username", api.vd="len($) > 0 && len($) <= 32")
    2: required string password (api.query="password", api.vd="len($) > 0 && len($) <= 32")
}

struct CreateUserResponse {
    1: required i64 status_code
    2: required string status_msg
    3: required i64 user_id
    4: required string token
}

struct QueryUserRequest {
    1: required string user_id (api.query="user_id") // TODO: string ?
    2: required string token (api.query="token")
}

struct QueryUserResponse {
    1: required i64 status_code
    2: required string status_msg
    3: required User user
}

service ApiService {
    DouyinUserLoginResponse UserLogin(1: DouyinUserLoginRequest req) (api.post="/douyin/user/login/")
    CreateUserResponse CreateUser(1: CreateUserRequest req) (api.post="/douyin/user/register/")
    QueryUserResponse QueryUser(1: QueryUserRequest req) (api.get="/douyin/user/")
}