package service

import (
	"douyin/common"
	"douyin/dao"
	"douyin/huancun"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

var UserService *UserServiceImpl
var userOnce sync.Once

func NewUserServiceInstance() *UserServiceImpl {
	userOnce.Do(
		func() {
			UserService = &UserServiceImpl{}
		})
	return UserService
}

type UserServiceImpl struct {
}

func (f *UserServiceImpl) Register(req common.RegisterReq) common.UserRegisterResp {
	username := req.Username
	password := req.Password
	user, err := dao.NewUserDaoInstance().QueryUserByName(username)
	resp := common.UserRegisterResp{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "用户已存在"
		return resp
	}
	enPWD, err := f.encodePWD(password)

	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "加密失败"
		return resp
	}

	id, err := f.register(username, enPWD)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "注册失败"
		return resp
	}
	resp.UserId = id
	resp.Token = f.tokenGenerate(username, enPWD)
	huancun.UsersLoginInfo[resp.Token] = &common.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      true,
	}
	return resp
}

func (f *UserServiceImpl) Login(req common.UserLoginReq) common.UserLoginResp {
	user, err := dao.NewUserDaoInstance().QueryUserByName(req.Username)
	rep := common.UserLoginResp{}
	if err != nil {
		rep.StatusCode = 1
		rep.StatusMsg = "用户或密码不正确"
		return rep
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		rep.StatusCode = 1
		rep.StatusMsg = "用户或密码不正确"
		return rep
	}
	token := f.tokenGenerate(user.Name, user.Password)
	rep.Token = token
	rep.UserId = user.Id
	huancun.UsersLoginInfo[rep.Token] = &common.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false,
	}
	return rep
}

func (f *UserServiceImpl) Info(req common.UserInfoReq) common.UserInfoResp {
	user, exist := huancun.UsersLoginInfo[req.Token]
	rep := common.UserInfoResp{}
	if exist {
		rep.User = *user
	} else {
		rep.StatusCode = 1
		rep.StatusMsg = "用户未登陆"
	}
	return rep
}

func (f *UserServiceImpl) register(username, enPWD string) (int64, error) {
	user := &dao.User{
		Name:     username,
		Password: enPWD,
	}
	if err := dao.NewUserDaoInstance().InsertUser(user); err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (f *UserServiceImpl) encodePWD(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (f *UserServiceImpl) tokenGenerate(username, enPWD string) string {
	return username + enPWD
}
