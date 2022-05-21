package controller

import (
	"douyin/common"
	"douyin/redis"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	//userId,_:=strconv.ParseInt(c.Query("user_id"),10,64)
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	commentText := c.Query("comment_text")
	commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	user, _ := redis.Get(token)
	if user == nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "用户未登陆",
		})
		return
	}
	req := common.CommentActionReq{
		UserId:      user.Id,
		Token:       token,
		VideoId:     videoId,
		ActionType:  int32(actionType),
		CommentText: commentText,
		CommentId:   commentId,
	}
	c.JSON(http.StatusOK, service.CommentService.Action(req))
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	token := c.Query("token")

	user, _ := redis.Get(token)
	if user == nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "用户未登陆",
		})
		return
	}
	req := common.CommentListReq{
		UserId:  user.Id,
		Token:   token,
		VideoId: videoId,
	}
	resp := service.CommentService.List(req)
	c.JSON(http.StatusOK, resp)
}
