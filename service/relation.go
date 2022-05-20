package service

import (
	"douyin/common"
	"douyin/dao"
	"sync"
)

var RelationService *RelationServiceImpl
var relationOnce sync.Once

func NewRelationServiceInstance() *RelationServiceImpl {
	relationOnce.Do(
		func() {
			RelationService = &RelationServiceImpl{}
		})
	return RelationService
}

type RelationServiceImpl struct {
}

func (f *RelationServiceImpl) Action(req common.RelationActionReq) common.RelationActionResp {
	followerId := req.FollowerId
	followeeId := req.FolloweeId
	actionType := req.ActionType
	resp := common.RelationActionResp{}
	if err := dao.RelationAction(followerId, followeeId, actionType); err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "关注错误"
	}
	return resp
}

func (f *RelationServiceImpl) FolloweeList(req common.FollowListReq) common.FollowListResp {
	userId := req.UserId
	relationList, err := dao.NewRelationDaoInstance().SelectFolloweeList(userId)
	resp := common.FollowListResp{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "关注列表出错"
		return resp
	}
	ul := make([]common.User, len(relationList))
	for i := range relationList {
		user, err := dao.NewUserDaoInstance().QueryUserById(relationList[i].ToUserId)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = "关注列表出错"
			return resp
		}
		ul[i].Id = user.Id
		ul[i].Name = user.Name
		ul[i].FollowCount = user.FollowCount
		ul[i].FollowerCount = user.FollowerCount
		ul[i].IsFollow = true
	}
	resp.UserList = ul
	return resp
}

func (f *RelationServiceImpl) FollowerList(req common.FollowListReq) common.FollowListResp {
	userId := req.UserId
	relationList, err := dao.NewRelationDaoInstance().SelectFollowerList(userId)
	resp := common.FollowListResp{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "粉丝列表出错"
		return resp
	}
	ul := make([]common.User, len(relationList))
	for i := range relationList {
		user, err := dao.NewUserDaoInstance().QueryUserById(relationList[i].UserId)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = "关注列表出错"
			return resp
		}
		ul[i].Id = user.Id
		ul[i].Name = user.Name
		ul[i].FollowCount = user.FollowCount
		ul[i].FollowerCount = user.FollowerCount
		isRelation, err := dao.NewRelationDaoInstance().IsRelation(userId, user.Id)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = "关注列表出错"
			return resp
		}
		if isRelation == 1 {
			ul[i].IsFollow = true
		}
	}
	resp.UserList = ul
	return resp
}
