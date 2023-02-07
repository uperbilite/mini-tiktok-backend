namespace go api.favorite

struct User {
    1: required i64 id
    2: required string name
    3: required i64 follow_count
    4: required i64 follower_count
    5: required bool is_follow
}

struct Video {
    1: required i64 id
    2: required User author
    3: required string play_url
    4: required string cover_url
    5: required i64 favorite_count
    6: required i64 comment_count
    7: required bool is_favorite
    8: required string title
}

struct DouyinFavoriteActionRequest {
    1: required string token (api.query="token")
    2: required i64 video_id (api.query="video_id")
    3: required i32 action_type (api.query="action_type" api.vd="in($, 1, 2)")
}

struct DouyinFavoriteActionResponse {
    1: required i32 status_code
    2: required string status_msg
}

struct DouyinFavoriteListRequest {
    1: required i64 user_id (api.query="user_id")
    2: required string token (api.query="token")
}

struct DouyinFavoriteListResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required list<Video> video_list
}

service ApiFavoriteService {
    DouyinFavoriteActionResponse DouyinFavoriteAction(1: DouyinFavoriteActionRequest req) (api.post="/douyin/favorite/action/")
    DouyinFavoriteListResponse DouyinFavoriteList(1: DouyinFavoriteListRequest req) (api.get="/douyin/favorite/list/")
}
