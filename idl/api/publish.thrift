namespace go api.publish

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

struct DouyinPublishActionRequest {
    1: required string token (api.form="token")
    2: required binary data (api.form="data")
    3: required string title (api.form="title")
}

struct DouyinPublishActionResponse {
    1: required i32 status_code
    2: required string status_msg
}

struct DouyinPublishListRequest {
    1: required string token (api.query="token")
    2: required i64 user_id (api.query="user_id")
}

struct DouyinPublishListResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required list<Video> video_list
}

service ApiPublishService {
    DouyinPublishActionResponse DouyinPublishAction(1: DouyinPublishActionRequest req) (api.post="/douyin/publish/action/")
    DouyinPublishListResponse DouyinPublishList(1: DouyinPublishListRequest req) (api.get="/douyin/publish/list/")
}
