package dao

import (
	"douyin/util"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Video struct {
	Id            int64     `gorm:"column:id"`
	UserId        int64     `gorm:"column:user_id"`
	PlayUrl       string    `gorm:"column:play_url"`
	CoverUrl      string    `gorm:"column:cover_url"`
	FavoriteCount int64     `gorm:"column:favorite_count"`
	CommentCount  int64     `gorm:"column:comment_count"`
	IsFavorite    bool      `gorm:"column:is_favorite"`
	CreateTime    time.Time `gorm:"column:create_time"`
}

func (Video) TableName() string {
	return "video"
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
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

func (v *VideoDao) SelectListByLimit(time int64, limit int) ([]*Video, error) {
	var video []*Video
	err := db.Where("unix_timestamp(create_time) <= ?", time).Order("create_time DESC").Limit(limit).Find(&video).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		util.Logger.Error("find video by userId err:" + err.Error())
		return nil, err
	}
	return video, nil
}

func (v *VideoDao) SelectById(id int64) ([]*Video, error) {
	var video []*Video
	err := db.Where("id = ?", id).Find(&video).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		util.Logger.Error("find video by userId err:" + err.Error())
		return nil, err
	}
	return video, nil
}

func (v *VideoDao) UpdateFavoriteById(userId int64, videoId int64, actionType int) error {
	db.Begin()
	var str string
	if actionType == 1 {
		str = "+"
	} else {
		str = "-"
	}
	if err := db.Table("video").Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count"+str+"?", 1)).Error; err != nil {
		db.Rollback()
		return err
	}
	if actionType == 1 {
		if err := NewFavoriteDaoInstance().CreateInstance(userId, videoId); err != nil {
			db.Rollback()
			return err
		}
	} else {
		if err := NewFavoriteDaoInstance().DeleteInstance(userId, videoId); err != nil {
			db.Rollback()
			return err
		}
	}
	db.Commit()
	return nil
}
