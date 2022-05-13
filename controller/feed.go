package controller

import (
	"douyin/common"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	latestTime := c.Query("latest_time")
	timeUnix, _ := strconv.ParseInt(latestTime, 10, 64)
	videoList, _, err := service.Feed(timeUnix, 30)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response:  common.Response{StatusCode: 1, StatusMsg: "视频流读取失败"},
			VideoList: []common.Video{},
			NextTime:  time.Now().Unix(),
		})
	}
	videoVOList := make([]common.Video, len(videoList))
	for i, videoInfo := range videoList {
		videoVOList[i].Id = videoInfo.Video.Id
		videoVOList[i].Author = common.User{
			Id:            videoInfo.User.Id,
			Name:          videoInfo.User.Name,
			FollowCount:   videoInfo.User.FollowCount,
			FollowerCount: videoInfo.User.FollowerCount,
			IsFollow:      videoInfo.User.IsFollow,
		}
		videoVOList[i].PlayUrl = videoInfo.Video.PlayUrl
		videoVOList[i].CoverUrl = videoInfo.Video.CoverUrl
		videoVOList[i].FavoriteCount = videoInfo.Video.FavoriteCount
		videoVOList[i].CommentCount = videoInfo.Video.CommentCount
		videoVOList[i].IsFavorite = videoInfo.Video.IsFavorite
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: videoVOList,
		NextTime:  time.Now().Unix(),
	})
}
