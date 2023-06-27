package main

import (
	"mcblog/config"
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

	r := gin.Default()

	routers.RouterInit(r)

	//fmt.Println(config.AppConfig)

	r.Run(config.AppConfig.Server.AppPort)

}
