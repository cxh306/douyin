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

func (f *FavoriteDao) SelectFavoriteByUserId(userId int64) ([]*Video, error) {
	var videos []*Video
	err := db.Table("favorite").
		Select("video.id as id, video.play_url as play_url,video.cover_url as cover_url,video.favorite_count as favorite_count,video.comment_count as comment_count,video.is_favorite as is_favorite").
		Joins("left join video on favorite.video_id = video.id where favorite.user_id=?", userId).Scan(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
func (f *FavoriteDao) CreateInstance(userId int64, videoId int64) error {
	if err := db.Model(Favorite{}).Create(map[string]interface{}{
		"user_id":  userId,
		"video_id": videoId,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (f *FavoriteDao) DeleteInstance(userId int64, videoId int64) error {
	if err := db.Where("user_id= ? and video_id=?", userId, videoId).Delete(&Favorite{}).Error; err != nil {
		return err
	}
	return nil
}
