namespace go comment

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

struct Comment {
    1: required i64 id
    2: required User user
    3: required string content
    4: required string create_date
}

struct CreateCommentRequest {
    1: required i64 user_id (vt.gt = "0")
    2: required i64 video_id (vt.gt = "0")
    3: required string content (vt.min_size = "1")
}

struct CreateCommentResponse {
    1: required BaseResp base_resp
    2: required Comment comment
}

struct DeleteCommentRequest {
    1: required i64 comment_id (vt.gt = "0")
}

struct DeleteCommentResponse {
    1: required BaseResp base_resp
}

struct GetCommentCountRequest {
    1: required i64 video_id (vt.gt = "0")
}

struct GetCommentCountResponse {
    1: required BaseResp base_resp
    2: required i64 comment_count
}

struct GetCommentListRequest {
    1: required i64 user_id (vt.gt = "0")
    2: required i64 video_id (vt.gt = "0")
}

struct GetCommentListResponse {
    1: required BaseResp base_resp
    2: required list<Comment> comment_list
}

service CommentService {
    CreateCommentResponse CreateComment(1: CreateCommentRequest req)
    DeleteCommentResponse DeleteComment(1: DeleteCommentRequest req)
    GetCommentCountResponse GetCommentCount(1: GetCommentCountRequest req)
    GetCommentListResponse GetCommentList(1: GetCommentListRequest req)
}