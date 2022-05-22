package dao

import "time"

func CommentAction(commentId int64, userId int64, videoId int64, actionType int32, commentText string) error {
	tx := db.Begin()
	comment := Comment{Id: commentId, UserId: userId, VideoId: videoId, CommentText: commentText, CreateTime: time.Now()}

	if actionType == 1 {
		if err := commentDao.CreateInstance(&comment); err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if err := commentDao.DeleteInstance(&comment); err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := videoDao.UpdateCommentCount(videoId, actionType); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func FavoriteAction(userId int64, videoId int64, actionType int32) error {
	tx := db.Begin()
	favorite := Favorite{UserId: userId, VideoId: videoId}

	if actionType == 1 {
		if err := favoriteDao.CreateInstance(favorite); err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if err := favoriteDao.DeleteInstance(favorite); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := videoDao.UpdateFavoriteById(videoId, actionType); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func RelationAction(followerId int64, followeeId int64, actionType int32) error {
	tx := db.Begin()
	if actionType == 1 {
		if err := relationDao.InsertRelation(Relation{UserId: followerId, ToUserId: followeeId}); err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if err := relationDao.DeleteRelation(Relation{UserId: followerId, ToUserId: followeeId}); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := userDao.UpdateFollowCnt(followerId, actionType); err != nil {
		tx.Rollback()
		return err
	}
	if err := userDao.UpdateFollowerCnt(followeeId, actionType); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
