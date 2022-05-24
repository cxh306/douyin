package dao

import "sync"

type Favorite struct {
	Id      int64 `gorm:"column:id"`
	UserId  int64 `gorm:"column:user_id"`
	VideoId int64 `gorm:"column:video_id"`
}

func (Favorite) TableName() string {
	return "favorite"
}

type FavoriteDao struct {
}

var favoriteDao *FavoriteDao
var favoriteOnce sync.Once

func NewFavoriteDaoInstance() *FavoriteDao {
	favoriteOnce.Do(
		func() {
			favoriteDao = &FavoriteDao{}
		})
	return favoriteDao
}

//func (*FavoriteDao) SelectFavoriteByUserId(userId int64) ([]*Video, error) {
//	var videos []*Video
//	err := db.Table("favorite").
//		Select("video.id as id, video.play_url as play_url,video.cover_url as cover_url,video.favorite_count as favorite_count,video.comment_count as comment_count,video.is_favorite as is_favorite").
//		Joins("join video on favorite.video_id = video.id where favorite.user_id=?", userId).Scan(&videos).Error
//	if err != nil {
//		return nil, err
//	}
//	return videos, nil
//}

func (*FavoriteDao) SelectFavoriteByUserId(userId int64) ([]*Favorite, error) {
	var favorites []*Favorite
	err := db.Where("user_id=?", userId).Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func (*FavoriteDao) CreateInstance(favorite Favorite) error {
	return db.Create(&favorite).Error
}

func (*FavoriteDao) DeleteInstance(favorite Favorite) error {
	return db.Where("user_id=? AND video_id=?", favorite.UserId, favorite.VideoId).Delete(&Favorite{}).Error
}

func (*FavoriteDao) IsFavorite(userId int64, videoId int64) (int64, error) {
	var count int64
	err := db.Model(&Favorite{}).Where("user_id=? AND video_id=?", userId, videoId).Count(&count).Error
	if err != nil {
		return -1, nil
	} else {
		return count, nil
	}
}
