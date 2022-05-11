package service

import "douyin/dao"

func FavoriteAction(userId int64, token string, videoId int64, actionType int) int {
	return NewFavoriteActionFlow(userId, token, videoId, actionType).Do()
}

type FavoriteActionFlow struct {
	UserId     int64
	Token      string
	VideoId    int64
	ActionType int

	Status int
}

func NewFavoriteActionFlow(userId int64, token string, videoId int64, actionType int) *FavoriteActionFlow {
	return &FavoriteActionFlow{UserId: userId, Token: token, VideoId: videoId, ActionType: actionType}
}
func (f *FavoriteActionFlow) Do() int {
	err := dao.NewVideoDaoInstance().UpdateFavoriteById(f.VideoId, f.ActionType)
	if err != nil {
		//更新失败
		f.Status = 1
	} else {
		//更新成功
		f.Status = 0
	}
	return f.Status
}
