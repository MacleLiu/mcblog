// config包提供参数配置相关的变量和方法定义
package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

func init() {
	InitializeConfig()
}

func InitializeConfig() {
	//配置文件路径
	configFile := "config/config.toml"
	//支持通过环境变量修改配置文件路径
	// if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
	// 	configFile = configEnv
	// }

	/*
		//初始化viper
		viper.SetConfigFile(configFile)
		viper.SetConfigType("toml")
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("read config failed: %s", err))
		}

		//将配置赋值给全局变量
		if err := viper.Unmarshal(&AppConfig); err != nil {
			log.Fatalf("config assignment failed: %v", err)
		}
	*/

	//使用viper解析到结构体失败，原因未知
	//所以暂时使用tmol的解析方法
	if _, err := toml.DecodeFile(configFile, &AppConfig); err != nil {
		panic(fmt.Errorf("read config failed: %s", err))
	}
}
