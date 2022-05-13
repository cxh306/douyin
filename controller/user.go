package controller

import (
	"douyin/common"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin

/**
默认登陆2个用户
*/
var usersLoginInfo = map[string]*common.User{
	"cxh$2a$10$AJuzrwWkhR.jUqzYQkt1seng.G5fDB./RJqrIQhmD0NCrxNDc6z1O": &common.User{
		Id:            1,
		Name:          "cxh",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	},
	"cxy$2a$10$CBgjkPMcKrxyOtQedzrYyuxw7Cu2tUfn8g6GZAyKCJ3TspFjH/rwO": &common.User{
		Id:            3,
		Name:          "cxy",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	},
}

type UserLoginResponse struct {
	common.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	common.Response
	User common.User `json:"user"`
}

func Register(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	id, token, err := service.Register(username, password)
	user, err := service.UserInfo(id)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		usersLoginInfo[token] = &common.User{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      false,
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 0},
			UserId:   id,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	id, token, err := service.Login(username, password)
	user, err := service.UserInfo(id)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "账号或密码错误"},
		})
	} else {
		usersLoginInfo[token] = &common.User{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      false,
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 0, StatusMsg: "success"},
			UserId:   id,
			Token:    token,
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 0},
			User:     *user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
