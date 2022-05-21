package controller

import (
	"douyin/common"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin

func Register(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")
	registerRequest := common.RegisterReq{
		Username: username,
		Password: password,
	}
	rep := service.UserService.Register(registerRequest)
	c.JSON(http.StatusOK, rep)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	req := common.UserLoginReq{Username: username, Password: password}
	resp := service.UserService.Login(req)
	c.JSON(http.StatusOK, resp)
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	req := common.UserInfoReq{UserId: userId, Token: token}
	rep := service.UserService.Info(req)
	c.JSON(http.StatusOK, rep)
}
