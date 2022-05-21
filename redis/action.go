package redis

import (
	"douyin/common"
	"encoding/json"
	"time"
)

func Set(token string, user common.User) error {
	userByte, _ := json.Marshal(user)
	return rdb.Set(token, string(userByte), time.Hour).Err()
}

func Get(token string) (*common.User, error) {
	user, _ := rdb.Get(token).Result()
	//用户未登陆
	if user == "" {
		return nil, nil
	}
	var res common.User
	err := json.Unmarshal([]byte(user), &res)
	return &res, err
}

func Delete(token string) (*common.User, error) {
	user, _ := rdb.Get(token).Result()
	//用户未登陆
	if user == "" {
		return nil, nil
	}
	var res common.User
	err := json.Unmarshal([]byte(user), &res)
	return &res, err
}
