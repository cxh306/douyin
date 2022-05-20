package controller

import (
	"douyin/common"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	latestTime := c.Query("latest_time")
	timeUnix, _ := strconv.ParseInt(latestTime, 10, 64)
	token := c.Query("token")
	req := common.FeedReq{
		LatestTime: timeUnix,
		Token:      token,
	}
	resp := service.VideoService.Feed(req)
	c.JSON(http.StatusOK, resp)
}
