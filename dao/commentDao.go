package dao

import (
	"sync"
	"time"
)

type Comment struct {
	Id          int64     `gorm:"column:id"`
	UserId      int64     `gorm:"column:user_id"`
	VideoId     int64     `gorm:"column:video_id"`
	CommentText string    `gorm:"column:comment_text"`
	CreateTime  time.Time `gorm:"column:create_time"`
}

func (Comment) TableName() string {
	return "comment"
}

type CommentDao struct {
}

var commentDao *CommentDao
var commentOnce sync.Once

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(
		func() {
			commentDao = &CommentDao{}
		})
	return commentDao
}

func (f *CommentDao) CreateInstance(userId int64, videoId int64, commentText string) error {
	if err := db.Model(Comment{}).Create(map[string]interface{}{
		"user_id":      userId,
		"video_id":     videoId,
		"comment_text": commentText,
		"comment_time": time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func (f *CommentDao) DeleteInstance(commentId int64) error {
	if err := db.Where("comment_id = ?", commentId).Delete(&Comment{}).Error; err != nil {
		return err
	}
	return nil
}

func (f *CommentDao) SelectCommentList(video int64) ([]*Comment, error) {
	var commentList []map[string]interface{}
	if err := db.Where("video_id = ?", video).Find(&commentList).Error; err != nil {
		return nil, err
	}
	return commentList, nil
}