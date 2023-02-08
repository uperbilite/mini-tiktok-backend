namespace go api.comment

struct User {
    1: required i64 id
    2: required string name
    3: required i64 follow_count
    4: required i64 follower_count
    5: required bool is_follow
}

struct Comment {
    1: required i64 id
    2: required User user
    3: required string content
    4: required string create_date
}

struct DouyinCommentActionRequest {
    1: required string token (api.query="token")
    2: required i64 video_id (api.query="video_id")
    3: required i32 action_type (api.query="action_type" api.vd="in($, 1, 2)")
    4: optional string comment_text (api.query="comment_text")
    5: optional i64 comment_id (api.query="comment_id")
}

struct DouyinCommentActionResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required Comment comment
}

struct DouyinCommentListRequest {
    1: required string token (api.query="token")
    2: required i64 video_id (api.query="video_id")
}

struct DouyinCommentListResponse {
    1: required i32 status_code
    2: required string status_msg
    3: required list<Comment> comment
}

service ApiCommentService {
    DouyinCommentActionResponse DouyinCommentAction(1: DouyinCommentActionRequest req) (api.post="/douyin/comment/action/")
    DouyinCommentListResponse DouyinCommentList(1: DouyinCommentListRequest req) (api.get="/douyin/comment/list/")
}