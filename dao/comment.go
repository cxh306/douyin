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

func (*CommentDao) DeleteInstance(comment *Comment) error {
	return db.Delete(comment).Error
}

func (*CommentDao) SelectComments(videoId int64) ([]Comment, error) {
	var commentList []Comment
	err := db.Where("video_id = ?", videoId).Find(&commentList).Error
	return commentList, err
}

func (*CommentDao) QueryComment(userId int64, videoId int64) (Comment, error) {
	var comment Comment
	err := db.Where("user_id = ? AND video_id=?", userId, videoId).Find(&comment).Error
	return comment, err
}
