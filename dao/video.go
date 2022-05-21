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
	Title         string    `gorm:"column:title"`
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

func (v *VideoDao) SelectById(id int64) (*Video, error) {
	var video *Video
	err := db.Where("id = ?", id).Find(&video).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return video, nil
}

func (v *VideoDao) SelectByIds(ids []int64) ([]*Video, error) {
	var video []*Video
	err := db.Where("id IN ?", ids).Find(&video).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return video, nil
}

func (v *VideoDao) UpdateFavoriteById(videoId int64, actionType int32) error {
	var str string
	if actionType == 1 {
		str = "+"
	} else {
		str = "-"
	}
	return db.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count"+str+"?", 1)).Error
}

func (v *VideoDao) InsertVideo(video Video) error {
	return db.Create(&video).Error
}

func (v *VideoDao) UpdateCommentCount(videoId int64, actionType int32) error {
	var str string
	if actionType == 1 {
		str = "+"
	} else {
		str = "-"
	}
	return db.Model(&Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count"+str+"?", 1)).Error
}

func (v *VideoDao) UpdateCoverUrl(videoId int64, actionType int32) error {
	var str string
	if actionType == 1 {
		str = "+"
	} else {
		str = "-"
	}
	return db.Model(&Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count"+str+"?", 1)).Error
}
