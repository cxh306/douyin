package dao

import (
	"gorm.io/gorm"
	"sync"
)

type User struct {
	Id            int64  `gorm:"column:id"`
	Name          string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
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
		return err
	}
	return nil
}

func (*UserDao) QueryUserByName(username string) (*User, error) {
	var user User
	err := db.Where("username=?", username).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
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
		return nil, err
	}
	return &user, nil
}

func (*UserDao) QueryUserByIds(ids []int64) ([]*User, error) {
	var user []*User
	err := db.Where("id IN ?", ids).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (*UserDao) UpdateFollowerCnt(userId int64, actionType int32) error {
	var str string
	if actionType == 1 {
		str = "+"
	} else {
		str = "-"
	}
	return db.Model(&User{}).Where("id = ?", userId).Update("follower_count", gorm.Expr("follower_count"+str+"?", 1)).Error
}

func (*UserDao) UpdateFollowCnt(userId int64, actionType int32) error {
	var str string
	if actionType == 1 {
		str = "+"
	} else {
		str = "-"
	}
	return db.Model(&User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count"+str+"?", 1)).Error
}
