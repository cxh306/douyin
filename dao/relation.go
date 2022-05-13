package dao

import (
	"douyin/common"
	"gorm.io/gorm"
	"sync"
)

type Relation struct {
	Id       int64 `gorm:"column:id"`
	UserId   int64 `gorm:"column:user_id"`
	toUserId int64 `gorm:"column:to_user_id"`
}

func (Relation) TableName() string {
	return "relation"
}

type RelationDao struct {
}

var relationDao *RelationDao
var relationOnce sync.Once

func NewRelationDaoInstance() *RelationDao {
	relationOnce.Do(
		func() {
			relationDao = &RelationDao{}
		})
	return relationDao
}

func (*RelationDao) UpdateRelation(userId int64, toUserId int64, actionType int) error {
	db.Begin()
	var str string
	if actionType == 1 {
		str = "+"
	} else {
		str = "-"
	}
	if actionType == 1 {
		if err := db.Table("relation").Model(Relation{}).Create(map[string]interface{}{
			"user_id":    userId,
			"to_user_Id": toUserId,
		}).Error; err != nil {
			db.Rollback()
			return err
		}
	} else {
		if err := db.Table("relation").Where("user_id = ? and to_user_id= ?", userId, toUserId).Delete(&Relation{}).Error; err != nil {
			db.Rollback()
			return err
		}
	}

	if err := db.Table("user").Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count"+str+"?", 1)).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Table("user").Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count"+str+"?", 1)).Error; err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

func (*RelationDao) SelectFollowList(userId int64) ([]common.User, error) {
	var followList []common.User
	err := db.Table("relation").Select("relation.to_user_id,user.name,user.follow_count,user.follower_count,user.is_follow").
		Joins("join user on relation.user_id =? and relation.to_user_id=user.id", userId).Find(&followList).Error
	return followList, err
}
