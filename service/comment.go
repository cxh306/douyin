package service

import "douyin/dao"

/**
CommentAction
*/
func CommentAction(id int64, userId int64, videoId int64, actionType int, commentText string) error {
	return NewCommentActionFlow(id, userId, videoId, actionType, commentText).Do()
}

type CommentActionFlow struct {
	Id          int64
	UserId      int64
	VideoId     int64
	ActionType  int
	CommentText string

	Status int
}

func NewCommentActionFlow(id int64, userId int64, videoId int64, actionType int, commentText string) *CommentActionFlow {
	return &CommentActionFlow{Id: id, UserId: userId, VideoId: videoId, ActionType: actionType, CommentText: commentText}
}
func (f *CommentActionFlow) Do() error {
	var err error
	if f.ActionType == 1 {
		err = dao.NewCommentDaoInstance().CreateInstance(f.UserId, f.VideoId, f.CommentText)
	} else {
		err = dao.NewCommentDaoInstance().DeleteInstance(f.Id)
	}
	if err != nil {
		//更新失败
		return err
	} else {
		//更新成功
		return nil
	}
}

/**
CommentList
*/

type User struct {
	Id            int64
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}

type Comment struct {
	Id         int64
	User       User
	Content    string
	CreateTime string
}

type CommentListFlow struct {
	VideoId int64

	CommentList []*Comment
}

func CommentList(videoId int64) ([]Comment, error) {
	return NewCommentListFlow(videoId).Do()
}

func NewCommentListFlow(videoId int64) *CommentListFlow {
	return &CommentListFlow{
		VideoId: videoId,
	}
}

func (f *CommentListFlow) Do() ([]Comment, error) {
	result, err := dao.NewCommentDaoInstance().SelectCommentList(f.VideoId)
	commentList := make([]Comment, len(result))
	for i, v := range result {
		commentList[i].Id = v.Id
		commentList[i].Content = v.CommentText
		commentList[i].CreateTime = v.CreateTime.Format("01-02")
		commentList[i].User = User{
			Id:            v.UserId,
			Name:          v.Name,
			FollowCount:   v.FollowCount,
			FollowerCount: v.FollowerCount,
			IsFollow:      v.IsFollow,
		}
	}
	if err != nil {
		return nil, err
	}
	return commentList, nil
}
