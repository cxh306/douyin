package service

import (
	"douyin/dao"
)

//type VideoModel struct {
//	Id            int64
//	UserId        int64
//	PlayUrl       string
//	CoverUrl      string
//	FavoriteCount int64
//	CommentCount  int64
//	IsFavorite    bool
//}

func NewPublishListFlow(userId int64) *PublishListFlow {
	return &PublishListFlow{userId: userId}
}

func PublisList(userId int64) ([]*dao.Video, error) {
	return NewPublishListFlow(userId).Do()
}

type PublishListFlow struct {
	userId int64

	videoList []*dao.Video
}

func (f *PublishListFlow) Do() ([]*dao.Video, error) {
	videoList, err := dao.NewVideoDaoInstance().SelectListByUserId(f.userId)
	if err != nil {
		return nil, err
	}
	f.videoList = videoList
	return f.videoList, nil
}

type PublishFlow struct {
	userId int64
}
