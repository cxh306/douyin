package service

import "douyin/dao"

/**
favoriteAction
*/
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

func FavoriteList(userId int64, token string) ([]*dao.Video, error) {
	return NewFavoriteListFlow(userId, token).Do()
}

type FavoriteListFlow struct {
	UserId int64
	Token  string

	VideoInFoList []*VideoInFo
}

func NewFavoriteListFlow(userId int64, token string) *FavoriteListFlow {
	return &FavoriteListFlow{
		UserId: userId,
		Token:  token,
	}
}

func (f *FavoriteListFlow) Do() ([]*dao.Video, error) {
	favorites, err := dao.NewFavoriteDaoInstance().SelectFavoriteByUserId(f.UserId)
	if err != nil {
		return nil, err
	}
	return favorites, nil
}
