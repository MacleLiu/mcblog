package main

import (
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
	modles.InitDb()

	gin.SetMode(config.AppConfig.Server.AppModel)

	r := gin.New()
	r.Use(middlewares.Logger(), middlewares.Cors())
	r.Use(gin.Recovery())

	routers.RouterInit(r)

	//fmt.Println(config.AppConfig)

	r.Run(config.AppConfig.Server.AppPort)

}
