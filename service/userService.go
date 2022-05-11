package service

import (
	"douyin/dao"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type UserModel struct {
	Id            int64
	Name          string
	FollowCount   int64
	FollowerCount int64

	Token string
}

type UserService struct {
}

var userOnce sync.Once
var userService *UserService

func NewUserService() *UserService {
	userOnce.Do(func() {
		userService = &UserService{}
	})
	return userService
}

func (f *UserService) Register(username, password string) (*UserModel, error) {
	encodePWD, err := f.encodePWD(password)
	if err != nil {
		return nil, err
	}
	id, err := f.register(username, encodePWD)
	if err != nil {
		return nil, err
	}

	return &UserModel{Id: id, Token: f.tokenGenerate(username, encodePWD)}, nil
}

func (f *UserService) Login(username, password string) (*UserModel, error) {
	user, err := dao.NewUserDaoInstance().QueryUserByName(username)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	token := f.tokenGenerate(user.Name, user.Password)

	return &UserModel{Id: user.Id, Token: token, Name: user.Name, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount}, nil
}

func (f *UserService) FindById(id int64) (*UserModel, error) {
	user, err := dao.NewUserDaoInstance().QueryUserById(id)
	if err != nil {
		return nil, err
	}
	return &UserModel{
		Id: user.Id, Name: user.Name,
		FollowCount:   user.FollowerCount,
		FollowerCount: user.FollowerCount,
	}, nil
}

func (f *UserService) encodePWD(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (f *UserService) register(username, encodePWD string) (int64, error) {
	user := &dao.User{
		Name:     username,
		Password: encodePWD,
	}
	if err := dao.NewUserDaoInstance().InsertUser(user); err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (f *UserService) tokenGenerate(username, encodePWD string) string {
	return username + encodePWD
}
