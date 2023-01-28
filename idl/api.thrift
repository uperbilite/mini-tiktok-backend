namespace go api

struct CheckUserRequest {
    1: required string username (api.query="username")
    2: required string password (api.query="password")
}

struct CheckUserResponse {
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

service ApiService {
    CheckUserResponse CheckUser(1: CheckUserRequest req) (api.post="/douyin/user/login/")
    CreateUserResponse CreateUser(1: CreateUserRequest req) (api.post="/douyin/user/register/")
}