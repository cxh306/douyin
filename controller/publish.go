package controller

import (
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	Response
	VideoList []VideoVO `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	user, exist := usersLoginInfo[token]
	if !exist {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "用户未登陆",
		})
		return
	}
	videoList, errVideo := service.PublisList(user.Id)
	if errVideo != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "视频错误",
		})
		return
	}
	videoVOList := make([]VideoVO, len(videoList))
	userVO := *user
	for i, video := range videoList {
		videoVOList[i] = VideoVO{
			Id:            video.Id,
			Author:        userVO,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
		}
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videoVOList,
	})
}
