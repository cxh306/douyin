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

func (*CommentDao) CreateInstance(comment *Comment) error {
	return db.Create(comment).Error
}

func (*CommentDao) DeleteInstance(commentId int64) error {
	return db.Where("comment_id = ?", commentId).Delete(&Comment{}).Error
}

func (*CommentDao) SelectCommentList(videoId int64) ([]Comment, error) {
	var commentList []Comment
	err := db.Where("video_id = ?", videoId).Find(&commentList).Error
	return commentList, err
}
