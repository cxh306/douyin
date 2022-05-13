package service

import (
	"douyin/common"
	"douyin/dao"
	"golang.org/x/crypto/bcrypt"
)

func NewRegisterFlow(username, password string) *RegisterFlow {
	return &RegisterFlow{username: username, password: password}
}

func NewLoginFlow(username, password string) *LoginFlow {
	return &LoginFlow{username: username, password: password}
}

func NewUserInfoFlow(id int64) *UserInfoFlow {
	return &UserInfoFlow{id: id}
}

func Register(username, password string) (int64, string, error) {
	return NewRegisterFlow(username, password).Do()
}

func Login(username, password string) (int64, string, error) {
	return NewLoginFlow(username, password).Do()
}

func UserInfo(id int64) (*common.User, error) {
	return NewUserInfoFlow(id).Do()
}

/**
Register
*/

type RegisterFlow struct {
	username, password string

	enPWD string

	userId int64
	token  string
}

func (f *RegisterFlow) Do() (int64, string, error) {
	if err := f.encodePWD(); err != nil {
		return 0, "", err
	}

	if err := f.register(); err != nil {
		return 0, "", err
	}
	f.tokenGenerate()
	return f.userId, f.token, nil
}

func (f *RegisterFlow) register() error {
	user := &dao.User{
		Name:     f.username,
		Password: f.enPWD,
	}
	if err := dao.NewUserDaoInstance().InsertUser(user); err != nil {
		return err
	}
	f.userId = user.Id
	return nil
}

func (f *RegisterFlow) encodePWD() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(f.password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return err
	}
	f.enPWD = string(hash)
	return nil
}

func (f *RegisterFlow) tokenGenerate() {
	f.token = f.username + f.enPWD
}

/**
Login
*/

type LoginFlow struct {
	username, password string

	enPWD string

	userId int64
	token  string
}

func (f *LoginFlow) Do() (int64, string, error) {
	user, err := dao.NewUserDaoInstance().QueryUserByName(f.username)
	if err != nil {
		return 0, "", err
	}
	f.enPWD = user.Password
	if err := bcrypt.CompareHashAndPassword([]byte(f.enPWD), []byte(f.password)); err != nil {
		return 0, "", err
	}
	f.userId = user.Id
	f.token = f.tokenGenerate()

	return f.userId, f.token, nil
}

func (f *LoginFlow) tokenGenerate() string {
	return f.username + f.enPWD
}

/**
UserInfo
*/
type UserInfoFlow struct {
	id int64

	user common.User
}

func (f *UserInfoFlow) Do() (*common.User, error) {
	user, err := dao.NewUserDaoInstance().QueryUserById(f.id)
	if err != nil {
		return nil, err
	}
	return &common.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowerCount,
		FollowerCount: user.FollowerCount,
	}, nil
}
