package service

import (
	"douyin/common"
	"douyin/dao"
	"sync"
)

var CommentService *CommentServiceImpl
var commentOnce sync.Once

func NewCommentServiceInstance() *CommentServiceImpl {
	favoriteOnce.Do(
		func() {
			CommentService = &CommentServiceImpl{}
		})
	return CommentService
}

type CommentServiceImpl struct {
}

func (f *CommentServiceImpl) Action(req common.CommentActionReq) common.CommentActionResp {
	resp := common.CommentActionResp{}
	if err := dao.CommentAction(req.UserId, req.VideoId, req.ActionType, req.CommentText); err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "评论错误"
	}
	return resp
}

func (f *CommentServiceImpl) List(req common.CommentListReq) common.CommentListResp {
	resp := common.CommentListResp{}
	videoId := req.VideoId
	cl, err := dao.NewCommentDaoInstance().SelectCommentList(videoId)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "评论列表出错"
		return resp
	}
	comments := make([]common.Comment, len(cl))
	for i := range cl {
		user, err := dao.NewUserDaoInstance().QueryUserById(cl[i].UserId)
		if err != nil {
			resp.StatusCode = 1
			resp.StatusMsg = "评论列表出错"
			return resp
		}

		author := common.User{}
		author.Id = user.Id
		author.Name = user.Name
		author.FollowCount = user.FollowCount
		author.FollowerCount = user.FollowerCount
		isRelation, err1 := dao.NewRelationDaoInstance().IsRelation(req.UserId, user.Id)
		if err1 != nil {
			resp.StatusCode = 1
			resp.StatusMsg = "评论列表出错"
			return resp
		}
		if isRelation == 1 || user.Id == req.UserId {
			author.IsFollow = true
		}
		comments[i].Id = cl[i].Id
		comments[i].User = author
		comments[i].Content = cl[i].CommentText
		comments[i].CreateDate = cl[i].CreateTime.Format("01-02")
	}
	resp.CommentList = comments
	return resp
}
