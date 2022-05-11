package dao

import (
	"douyin/util"
	"gorm.io/gorm"
	"sync"
)

type User struct {
	Id            int64  `gorm:"column:id"`
	Name          string `gorm:"column:name"`
	Password      string `gorm:"column:password"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
	IsFollow      bool   `gorm:"column:is_follow"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) InsertUser(user *User) error {
	if err := db.Create(user).Error; err != nil {
		util.Logger.Error("insert post err:" + err.Error())
		return err
	}
	return nil
}

func (*UserDao) QueryUserByName(username string) (*User, error) {
	var user User
	err := db.Where("name=?", username).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find user by name err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

func (*UserDao) QueryUserById(id int64) (*User, error) {
	var user User
	err := db.Where("id=?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}
	return &user, nil
}
