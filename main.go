package main

import (
	"douyin/dao"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := Init(); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.Use(gin.Logger())

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Init() error {
	if err := dao.Init(); err != nil {
		return err
	}
	if err := util.InitLogger(); err != nil {
		return err
	}
	dao.NewUserDaoInstance()
	dao.NewVideoDaoInstance()
	dao.NewVideoDaoInstance()
	return nil
}
