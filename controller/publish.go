package controller

import (
	"bytes"
	"douyin/common"
	"douyin/redis"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token, _ := c.GetPostForm("token")
	title, _ := c.GetPostForm("title")
	user, _ := redis.Get(token)
	if user == nil {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "用户未登陆"})
		return
	}
	file, _, err := c.Request.FormFile("data")
	defer file.Close()
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "视频上传错误",
		})
		return
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "视频上传错误",
		})
		return
	}
	req := common.PublishReq{Token: token, Data: buf.Bytes(), Title: title}
	resp := service.VideoService.Publish(req)
	c.JSON(http.StatusOK, resp)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	user, _ := redis.Get(token)
	if user == nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "用户未登陆",
		})
		return
	}
	req := common.PublishListReq{UserId: user.Id, Token: token}
	resp := service.VideoService.PublishList(req)
	c.JSON(http.StatusOK, resp)
}
