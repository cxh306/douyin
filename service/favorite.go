package service

import "douyin/dao"

/**
favoriteAction
*/
func FavoriteAction(userId int64, videoId int64, actionType int) int {
	return NewFavoriteActionFlow(userId, videoId, actionType).Do()
}

type FavoriteActionFlow struct {
	UserId     int64
	VideoId    int64
	ActionType int

	Status int
}

func NewFavoriteActionFlow(userId int64, videoId int64, actionType int) *FavoriteActionFlow {
	return &FavoriteActionFlow{UserId: userId, VideoId: videoId, ActionType: actionType}
}
func (f *FavoriteActionFlow) Do() int {
	err := dao.NewVideoDaoInstance().UpdateFavoriteById(f.UserId, f.VideoId, f.ActionType)
	if err != nil {
		//更新失败
		f.Status = 1
	} else {
		//更新成功
		f.Status = 0
	}
	return f.Status
}

/**
FavoriteList
*/

func FavoriteList(userId int64) ([]*dao.Video, error) {
	return NewFavoriteListFlow(userId).Do()
}

type FavoriteListFlow struct {
	UserId int64

	VideoInFoList []*VideoInFo
}

func NewFavoriteListFlow(userId int64) *FavoriteListFlow {
	return &FavoriteListFlow{
		UserId: userId,
	}
}

func (f *FavoriteListFlow) Do() ([]*dao.Video, error) {
	favorites, err := dao.NewFavoriteDaoInstance().SelectFavoriteByUserId(f.UserId)
	if err != nil {
		return nil, err
	}
	return favorites, nil
}
