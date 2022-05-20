package controller

import (
	"douyin/common"
	"douyin/huancun"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	//followerId,_ := strconv.ParseInt(c.Query("user_id"),10,64)
	followeeId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	user, exist := huancun.UsersLoginInfo[token]
	if !exist {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "用户未登陆",
		})
		return
	}

	//if followerId != user.Id{
	//	c.JSON(http.StatusOK,common.Response{
	//		StatusCode: 1,
	//		StatusMsg: "请求非法",
	//	})
	//	return
	//}
	req := common.RelationActionReq{}
	req.FollowerId = user.Id
	req.FolloweeId = followeeId
	req.ActionType = int32(actionType)
	req.Token = token
	c.JSON(http.StatusOK, service.NewRelationServiceInstance().Action(req))
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	//userId,_ := strconv.ParseInt(c.Query("user_id"),10,64)
	user, exist := huancun.UsersLoginInfo[token]
	if !exist {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "关注列表错误",
		})
		return
	}

	//if userId!=user.Id {
	//	c.JSON(http.StatusOK,common.Response{
	//		StatusCode: 1,
	//		StatusMsg: "请求非法",
	//	})
	//	return
	//}
	req := common.FollowListReq{
		UserId: user.Id,
		Token:  token,
	}
	c.JSON(http.StatusOK, service.NewRelationServiceInstance().FolloweeList(req))
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	//userId,_ := strconv.ParseInt(c.Query("user_id"),10,64)
	user, exist := huancun.UsersLoginInfo[token]
	if !exist {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "关注列表错误",
		})
		return
	}

	//if userId!=user.Id {
	//	c.JSON(http.StatusOK,common.Response{
	//		StatusCode: 1,
	//		StatusMsg: "请求非法",
	//	})
	//	return
	//}
	req := common.FollowListReq{
		UserId: user.Id,
		Token:  token,
	}
	c.JSON(http.StatusOK, service.NewRelationServiceInstance().FollowerList(req))
}
