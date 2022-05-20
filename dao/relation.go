package dao

import (
	"sync"
)

type Relation struct {
	Id       int64 `gorm:"column:id"`
	UserId   int64 `gorm:"column:follower_id"`
	ToUserId int64 `gorm:"column:followee_id"`
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

func (*RelationDao) DeleteRelation(relation Relation) error {
	return db.Delete(&relation).Error
}

func (*RelationDao) InsertRelation(relation Relation) error {
	return db.Create(&relation).Error
}

func (*RelationDao) SelectFolloweeList(userId int64) ([]*Relation, error) {
	var followeeList []*Relation
	err := db.Where("follower_id=?", userId).Find(&followeeList).Error
	return followeeList, err
}

func (*RelationDao) SelectFollowerList(userId int64) ([]*Relation, error) {
	var followerList []*Relation
	err := db.Where("followee_id=?", userId).Find(&followerList).Error
	return followerList, err
}

func (*RelationDao) IsRelation(followerId int64, followeeId int64) (int64, error) {
	var count int64
	err := db.Model(&Relation{}).Where("follower_id=? and followee_id=?", followerId, followeeId).Count(&count).Error
	if err != nil {
		return -1, nil
	} else {
		return count, nil
	}
}
