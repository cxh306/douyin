package service

import (
	"douyin/common"
	"douyin/dao"
	"sync"
)

/**
favoriteAction
*/

var FavoriteService *FavoriteServiceImpl
var favoriteOnce sync.Once

func NewFavoriteInstance() *FavoriteServiceImpl {
	favoriteOnce.Do(
		func() {
			FavoriteService = &FavoriteServiceImpl{}
		})
	return FavoriteService
}

type FavoriteServiceImpl struct {
}

func (f *FavoriteServiceImpl) Action(req common.FavoriteActionReq) common.FavoriteActionResp {
	userId := req.UserId
	videoId := req.VideoId
	actionType := req.ActionType
	resp := common.FavoriteActionResp{}
	if err := dao.FavoriteAction(userId, videoId, actionType); err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "点赞出错"
	}
	return resp
}

func (f *FavoriteServiceImpl) List(req common.FavoriteListReq) common.PublishListResp {
	resp := common.PublishListResp{}
	userId := req.UserId
	favorites, err := dao.NewFavoriteDaoInstance().SelectFavoriteByUserId(userId)
	videoIds := make([]int64, len(favorites))
	for i := range videoIds {
		videoIds[i] = favorites[i].VideoId
	}
	videoList, err := dao.NewVideoDaoInstance().SelectByIds(videoIds)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "点赞列表出错"
		return resp
	}
	v := make([]common.Video, len(videoList))
	for i := range v {
		v[i].Id = videoList[i].Id
		v[i].PlayUrl = videoList[i].PlayUrl
		v[i].CoverUrl = videoList[i].CoverUrl
		v[i].FavoriteCount = videoList[i].FavoriteCount
		v[i].CommentCount = videoList[i].CommentCount
		v[i].IsFavorite = true
		v[i].Title = videoList[i].Title
		author := common.User{}
		user, err := dao.NewUserDaoInstance().QueryUserById(videoList[i].UserId)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = "点赞列表出错"
			return resp
		}
		author.Id = user.Id
		author.Name = user.Name
		author.FollowCount = user.FollowCount
		author.FollowerCount = user.FollowerCount
		isRelation, err := dao.NewRelationDaoInstance().IsRelation(userId, author.Id)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = "点赞列表出错"
			return resp
		}
		if isRelation == 1 {
			author.IsFollow = true
		} else {
			author.IsFollow = false
		}
		v[i].Author = author
	}
	resp.VideoList = v
	return resp
}
