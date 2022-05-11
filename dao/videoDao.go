package dao

import (
	"douyin/util"
	"gorm.io/gorm"
	"sync"
)

type Video struct {
	Id            int64  `gorm:"column:id"`
	UserId        int64  `gorm:"column:user_id"`
	PlayUrl       string `gorm:"column:play_url"`
	CoverUrl      string `gorm:"column:cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	IsFavorite    bool   `gorm:"column:is_favorite"`
}

func (Video) TableName() string {
	return "video"
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	userOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

func (v *VideoDao) SelectListByUserId(userId int64) ([]*Video, error) {
	var video []*Video
	err := db.Where("user_id=?", userId).Find(&video).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		util.Logger.Error("find video by userId err:" + err.Error())
		return nil, err
	}
	return video, nil
}

func (v *VideoDao) SelectList() ([]*Video, error) {
	var video []*Video
	err := db.Find(&video).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		util.Logger.Error("find video by userId err:" + err.Error())
		return nil, err
	}
	return video, nil
}
