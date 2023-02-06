namespace go api.feed

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

struct DouyinFeedRequest {
    1: optional i64 latest_time (api.query="latest_time")
    2: optional string token (api.query="token")
}

struct DouyinFeedResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required i64 next_time
    4: required list<Video> video_list
}

service ApiFeedService {
    DouyinFeedResponse DouyinFeed(1: DouyinFeedRequest req) (api.get="/douyin/feed/")
}