package controller

import (
	"douyin/common"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	common.Response
	CommentList []common.Comment `json:"comment_list,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	token := c.Query("token")
	var userId int64
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.Atoi(c.Query("action_type"))
	commentText := c.Query("comment_text")

	if _, exist := usersLoginInfo[token]; exist {
		userId = usersLoginInfo[token].Id
		if err := service.CommentAction(id, userId, videoId, actionType, commentText); err == nil {
			c.JSON(http.StatusOK, common.Response{StatusCode: 0})
		} else {
			c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "评论更新失败"})
		}
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "用户不存在"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		result, err := service.CommentList(videoId)
		if err != nil {
			c.JSON(http.StatusOK, common.Response{StatusCode: 1})
		} else {
			commentList := make([]common.Comment, len(result))
			for i, v := range result {
				commentList[i].Id = v.Id
				commentList[i].Content = v.Content
				commentList[i].CreateDate = v.CreateTime
				commentList[i].User = common.User{
					Id:            v.User.Id,
					Name:          v.User.Name,
					FollowCount:   v.User.FollowCount,
					FollowerCount: v.User.FollowerCount,
					IsFollow:      v.User.IsFollow,
				}
			}
			c.JSON(http.StatusOK, CommentListResponse{
				Response:    common.Response{StatusCode: 0},
				CommentList: commentList,
			})
		}
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "用户不存在"})
	}
}
