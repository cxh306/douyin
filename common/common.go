package common

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type User struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type Video struct {
	Id            int64  `json:"id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type Comment struct {
	Id         int64  `json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type RegisterReq struct {
	Username string
	Password string
}

type UserRegisterResp struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserLoginReq struct {
	Username string
	Password string
}

type UserLoginResp struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserInfoReq struct {
	UserId int64
	Token  string
}

type UserInfoResp struct {
	Response
	User User `json:"user"`
}

type FeedReq struct {
	LatestTime int64
	Token      string
}

type FeedResp struct {
	Response
	VideoList []Video `json:"video_list"`
	NextTime  int64   `json:"next_time"`
}

type PublishReq struct {
	Token string
	Data  []byte
	Title string
}

type PublishResp struct {
	Response
}

type PublishListReq struct {
	UserId int64
	Token  string
}

type PublishListResp struct {
	Response
	VideoList []Video `json:"video_list"`
}

type FavoriteActionReq struct {
	UserId     int64
	Token      string
	VideoId    int64
	ActionType int32
}

type FavoriteActionResp struct {
	Response
}

type FavoriteListReq struct {
	UserId int64
	Token  string
}

type FavoriteListResp struct {
	Response
	VideoList []Video `json:"video_list"`
}

type CommentActionReq struct {
	UserId      int64
	Token       string
	VideoId     int64
	ActionType  int32
	CommentText string
	CommentId   int64
}

type CommentActionResp struct {
	Response
}

type CommentListReq struct {
	UserId  int64
	Token   string
	VideoId int64
}

type CommentListResp struct {
	Response
	CommentList []Comment `json:"comment_list"`
}

type RelationActionReq struct {
	FollowerId int64
	Token      string
	FolloweeId int64
	ActionType int32
}

type RelationActionResp struct {
	Response
}

type FollowListReq struct {
	UserId int64
	Token  string
}

type FollowListResp struct {
	Response
	UserList []User `json:"user_list"`
}

type FollowerListReq struct {
	UserId int64
	Token  string
}

type FollowerListResp struct {
	Response
	UserList []User `json:"user_list"`
}
