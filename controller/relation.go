package controller

import (
	"douyin/common"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	common.Response
	UserList []common.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.Atoi(c.Query("action_type"))
	if user, exist := usersLoginInfo[token]; exist {
		err := service.RelationAction(user.Id, toUserId, actionType)
		if err != nil {
			c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "关注出错"})
		}
		c.JSON(http.StatusOK, common.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	if user, exist := usersLoginInfo[token]; exist {
		followList, err := service.FollowList(user.Id)
		if err != nil {
			c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "关注列表出错"})
		} else {
			c.JSON(http.StatusOK, UserListResponse{
				Response: common.Response{
					StatusCode: 0,
				},
				UserList: followList,
			})
		}
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		UserList: []common.User{},
	})
}
