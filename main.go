package main

import (
	"fmt"
	"mcblog/config"
	"mcblog/middlewares"
	"mcblog/modles"
	"mcblog/routers"

	"github.com/gin-gonic/gin"
)

type Config struct {
	AppModel string `toml:"app_model"`
	AppPort  string `toml:"app_port"`
}

func main() {
	//初始化数据库
	fmt.Println("数据库初始化")
	modles.InitDb()
	fmt.Println("数据库初始化完成")
	gin.SetMode(config.AppConfig.Server.AppModel)

	r := gin.New()
	r.Use(middlewares.Logger(), middlewares.Cors())
	r.Use(gin.Recovery())

	routers.RouterInit(r)

	//fmt.Println(config.AppConfig)

	fmt.Println("服务启动完成")
	r.Run(config.AppConfig.Server.AppPort)

}
