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

type Result struct {
	Id            int64     `json:"id"`
	CommentText   string    `json:"comment_text"`
	CreateTime    time.Time `json:"create_time"`
	UserId        int64     `json:"user_id"`
	Name          string    `json:"name"`
	FollowCount   int64     `json:"follow_count"`
	FollowerCount int64     `json:"follower_count"`
	IsFollow      bool      `json:"is_follow"`
}

func (f *CommentDao) SelectCommentList(videoId int64) ([]Result, error) {
	var commentList []Result
	err := db.Table("comment").
		Select("comment.id, comment.comment_text,comment.create_time,user.id as user_id,user.name,user.follow_count,user.follower_count,user.is_follow").
		Joins("left join user on comment.user_id = user.id and comment.video_id=?", videoId).Scan(&commentList).Error
	if err != nil {
		return nil, err
	}
	return commentList, nil
}
