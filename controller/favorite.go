package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.Atoi(c.Query("action_type"))
	status := service.FavoriteAction(userId, token, videoId, actionType)
	if _, exist := usersLoginInfo[token]; exist && status == 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: []VideoVO{},
	})
}
