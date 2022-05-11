package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "用户未登陆"})
		return
	}
	userId := usersLoginInfo[token].Id
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.Atoi(c.Query("action_type"))
	status := service.FavoriteAction(userId, token, videoId, actionType)
	if status != 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "点赞出错"})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	user, exist := usersLoginInfo[token]
	favorites, err := service.FavoriteList(user.Id, token)

	videoList := make([]VideoVO, len(favorites))
	for i, v := range favorites {
		videoList[i].Id = v.Id
		videoList[i].Author = UserVO{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		}
		videoList[i].PlayUrl = v.PlayUrl
		videoList[i].CoverUrl = v.CoverUrl
		videoList[i].FavoriteCount = v.FavoriteCount
		videoList[i].CommentCount = v.CommentCount
		videoList[i].IsFavorite = v.IsFavorite
	}
	if exist && err == nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: videoList,
		})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "点赞视频列表出错",
			},
			VideoList: []VideoVO{},
		})
	}
}
