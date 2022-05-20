package controller

import (
	"douyin/common"
	"douyin/huancun"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	//userId,_ := strconv.ParseInt(c.Query("user_id"),10,64)
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	user, exist := huancun.UsersLoginInfo[token]
	if !exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "用户未登陆"})
		return
	}

	req := common.FavoriteActionReq{
		UserId:     user.Id,
		Token:      token,
		VideoId:    videoId,
		ActionType: int32(actionType),
	}

	resp := service.FavoriteService.Action(req)
	c.JSON(http.StatusOK, resp)
}

func FavoriteList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	user, exist := huancun.UsersLoginInfo[token]
	if !exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "用户未登陆"})
	}
	if userId != user.Id {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "请求非法"})
	}
	req := common.FavoriteListReq{
		UserId: userId,
		Token:  token,
	}
	resp := service.FavoriteService.List(req)
	c.JSON(http.StatusOK, resp)
}
