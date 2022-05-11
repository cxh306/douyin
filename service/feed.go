package service

import (
	"douyin/dao"
)

type VideoInFo struct {
	Video dao.Video
	User  dao.User
}

func Feed(latestTime int64, limit int) ([]*VideoInFo, int64, error) {
	return NewFeedFlow(latestTime, limit).Do()
}

type FeedFlow struct {
	latestTime int64
	limit      int

	VideoInFoList []*VideoInFo
	nextTime      int64
}

func NewFeedFlow(latestTime int64, limit int) *FeedFlow {
	return &FeedFlow{latestTime: latestTime, limit: limit}
}
func (f *FeedFlow) Do() ([]*VideoInFo, int64, error) {
	videoList, err := dao.NewVideoDaoInstance().SelectListByLimit(f.latestTime, f.limit)
	if err != nil {
		return nil, 0, err
	}
	VideoInFoList := make([]*VideoInFo, len(videoList))
	for i, video := range videoList {
		user, err := dao.NewUserDaoInstance().QueryUserById(video.UserId)
		if err != nil {
			return nil, 0, err
		}
		VideoInFoList[i] = &VideoInFo{
			Video: *video,
			User: dao.User{
				Id:            user.Id,
				Name:          user.Name,
				FollowCount:   user.FollowerCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      false,
			},
		}
	}
	f.nextTime = videoList[len(videoList)-1].CreateTime.Unix()
	f.VideoInFoList = VideoInFoList
	return f.VideoInFoList, f.nextTime, nil
}
