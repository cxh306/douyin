package service

import (
	"douyin/dao"
	"sync"
)

type VideoModel struct {
	Id            int64
	UserId        int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool
}

type VideoService struct {
}

var videoOnce sync.Once
var videoService *VideoService

func NewVideoService() *VideoService {
	videoOnce.Do(func() {
		videoService = &VideoService{}
	})
	return videoService
}

func (v *VideoService) FindVideoListByUserId(userId int64) ([]*VideoModel, error) {
	videoList, err := dao.NewVideoDaoInstance().SelectListByUserId(userId)
	if err != nil {
		return nil, err
	}
	videoModelList := make([]*VideoModel, len(videoList))
	for i, video := range videoList {
		videoModelList[i] = &VideoModel{
			Id:            video.Id,
			UserId:        video.UserId,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
		}
	}

	return videoModelList, nil
}

func (v *VideoService) FindVideoList() ([]*VideoModel, error) {
	videoList, err := dao.NewVideoDaoInstance().SelectList()
	if err != nil {
		return nil, err
	}
	videoModelList := make([]*VideoModel, len(videoList))
	for i, video := range videoList {
		videoModelList[i] = &VideoModel{
			Id:            video.Id,
			UserId:        video.UserId,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
		}
	}

	return videoModelList, nil
}
