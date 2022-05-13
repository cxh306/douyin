package service

import (
	"douyin/common"
	"douyin/dao"
)

/**
CommentAction
*/
func RelationAction(userId int64, toUserId int64, actionType int) error {
	return NewRelationActionFlow(userId, toUserId, actionType).Do()
}

type RelationActionFlow struct {
	UserId     int64
	ToUserId   int64
	ActionType int
}

func NewRelationActionFlow(userId int64, toUserId int64, actionType int) *RelationActionFlow {
	return &RelationActionFlow{UserId: userId, ToUserId: toUserId, ActionType: actionType}
}
func (f *RelationActionFlow) Do() error {
	var err error
	err = dao.NewRelationDaoInstance().UpdateRelation(f.UserId, f.ToUserId, f.ActionType)
	if err != nil {
		//更新失败
		return err
	} else {
		//更新成功
		return nil
	}
}

/**
FollowList
*/

func FollowList(userId int64) ([]common.User, error) {
	return NewFollowListFlow(userId).Do()
}

type FollowListFlow struct {
	UserId int64

	followList []common.User
}

func NewFollowListFlow(userId int64) *FollowListFlow {
	return &FollowListFlow{UserId: userId}
}
func (f *FollowListFlow) Do() ([]common.User, error) {
	var err error
	f.followList, err = dao.NewRelationDaoInstance().SelectFollowList(f.UserId)
	return f.followList, err
}
