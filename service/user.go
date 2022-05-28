package service

import (
	"douyin/common"
	"douyin/dao"
	"douyin/redis"
	"github.com/dgrijalva/jwt-go"
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
	resp.Token = f.tokenGenerate(username, password)

	userResp := common.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      true,
	}
	token := resp.Token
	if err := redis.Set(token, userResp); err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
	}
	return resp
}

func (f *UserServiceImpl) Login(req common.UserLoginReq) common.UserLoginResp {
	user, err := dao.NewUserDaoInstance().QueryUserByName(req.Username)
	resp := common.UserLoginResp{}
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "用户或密码不正确"
		return resp
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "用户或密码不正确"
		return resp
	}
	username := user.Name
	password := req.Password
	token := f.tokenGenerate(username, password)

	userRedis, err := redis.Get(token)

	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		return resp
	}
	if userRedis != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "用户已登陆"
		return resp
	}

	resp.Token = token
	resp.UserId = user.Id
	userRes := common.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false,
	}
	if err := redis.Set(token, userRes); err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
	}
	return resp
}

func (f *UserServiceImpl) Info(req common.UserInfoReq) common.UserInfoResp {
	resp := common.UserInfoResp{}
	token := req.Token
	userRedis, err := redis.Get(token)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err.Error()
		return resp
	}
	if userRedis == nil {
		resp.StatusCode = 1
		resp.StatusMsg = "用户未登陆"
		return resp
	}
	resp.User = *userRedis
	return resp
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

func (f *UserServiceImpl) tokenGenerate(username, password string) string {
	dict := make(map[string]interface{})
	dict["username"] = username
	dict["password"] = password
	token, _ := GenerateToken(dict, "1a2b3c") // 生成token
	return token
}

// GenerateToken 生成Token值
func GenerateToken(mapClaims jwt.MapClaims, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return token.SignedString([]byte(key))
}
