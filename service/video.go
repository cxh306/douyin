package service

import (
	"douyin/common"
	"douyin/config"
	"douyin/dao"
	"douyin/huancun"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var VideoService *VideoServiceImpl
var videoOnce sync.Once

func NewVideoServiceInstance() *VideoServiceImpl {
	videoOnce.Do(
		func() {
			VideoService = &VideoServiceImpl{}
		})
	return VideoService
}

type VideoServiceImpl struct {
}

func (f *VideoServiceImpl) Feed(req common.FeedReq) common.FeedResp {
	latestTime := req.LatestTime
	token := req.Token
	resp := common.FeedResp{}
	user, isLogin := huancun.UsersLoginInfo[token]
	limit := 30
	videoList, err := dao.NewVideoDaoInstance().SelectListByLimit(latestTime, limit)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "视频流出错"
		return resp
	}
	v := make([]common.Video, len(videoList))

	for i := range videoList {
		user, err := dao.NewUserDaoInstance().QueryUserById(videoList[i].UserId)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = "视频流出错"
			return resp
		}
		author := common.User{}
		author.Id = user.Id
		author.Name = user.Name
		author.FollowCount = user.FollowCount
		author.FollowerCount = user.FollowerCount
		v[i].Author = author
		v[i].PlayUrl = videoList[i].PlayUrl
		v[i].Id = videoList[i].Id
		v[i].CoverUrl = videoList[i].CoverUrl
		v[i].FavoriteCount = videoList[i].FavoriteCount
		v[i].CommentCount = videoList[i].CommentCount
		v[i].Title = videoList[i].Title
	}
	if isLogin {
		for i := range v {
			isRelation, err1 := dao.NewRelationDaoInstance().IsRelation(user.Id, v[i].Author.Id)
			if err1 != nil {
				resp.StatusCode = 1
				resp.StatusMsg = "视频流出错"
				return resp
			}
			isFavorite, err2 := dao.NewFavoriteDaoInstance().IsFavorite(user.Id, v[i].Id)
			if err2 != nil {
				resp.StatusCode = 1
				resp.StatusMsg = "视频流出错"
				return resp
			}
			if isRelation == 1 {
				v[i].Author.IsFollow = true
			}
			if isFavorite == 1 {
				v[i].IsFavorite = true
			}
		}
	}
	resp.VideoList = v
	return resp
}

func (f *VideoServiceImpl) Publish(req common.PublishReq) common.PublishResp {
	resp := common.PublishResp{}
	token := req.Token
	data := req.Data
	title := req.Title
	user, _ := huancun.UsersLoginInfo[token]
	filename := fmt.Sprintf("%d_%s.mp4", user.Id, title)
	saveFile := filepath.Join("./public/", filename)

	if err := ioutil.WriteFile(saveFile, data, 0666); err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "视频保存出错"
		return resp
	}
	playUrl := config.Url + "/static/" + filename
	video := dao.Video{UserId: user.Id, PlayUrl: playUrl, Title: title, CreateTime: time.Now()}
	if err := dao.NewVideoDaoInstance().InsertVideo(video); err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "视频入库出错"
		err := os.Remove(saveFile)
		if err != nil {
			resp.StatusMsg += "&本地文件删除出错"
		}
	}
	return resp
}

func (f *VideoServiceImpl) PublishList(req common.PublishListReq) common.PublishListResp {
	resp := common.PublishListResp{}
	userId := req.UserId
	vd, err := dao.NewVideoDaoInstance().SelectListByUserId(userId)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "视频发布列表出错"
		return resp
	}
	vs := make([]common.Video, len(vd))
	for i := range vd {
		vs[i].Id = vd[i].Id
		vs[i].PlayUrl = vd[i].PlayUrl
		vs[i].CoverUrl = vd[i].CoverUrl
		vs[i].FavoriteCount = vd[i].FavoriteCount
		vs[i].CommentCount = vd[i].CommentCount
		vs[i].Title = vd[i].Title
		userId := vd[i].UserId
		user, err := dao.NewUserDaoInstance().QueryUserById(userId)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = "视频用户出错"
			return resp
		}
		author := common.User{}
		author.Id = user.Id
		author.Name = user.Name
		author.FollowCount = user.FollowCount
		author.FollowerCount = user.FollowerCount
		author.IsFollow = true
		vs[i].Author = author
	}
	resp.VideoList = vs
	return resp
}
