namespace go video

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

struct GetVideosRequest {
    1: required i64 user_id
    2: required list<i64> video_ids
}

struct GetVideosResponse {
    1: required BaseResp base_resp
    2: required list<Video> videos
}

service VideoService {
    GetVideosResponse GetVideos(1: GetVideosRequest req)
}