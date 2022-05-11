package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []VideoVO `json:"video_list,omitempty"`
	NextTime  int64     `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	videoModelList, err := service.NewVideoService().FindVideoList()
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 1, StatusMsg: "视频流读取失败"},
			VideoList: []VideoVO{},
			NextTime:  time.Now().Unix(),
		})
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: []VideoVO{},
		NextTime:  time.Now().Unix(),
	})
}
