package config

type Config struct {
	Server    Server      `toml:"server"`
	Mysql     MysqlConfig `toml:"database"`
	JwtSecret string      `toml:"jwtSecret"`
}

type Server struct {
	AppModel string `toml:"app_model"`
	AppPort  string `toml:"app_port"`
}

type MysqlConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DbName   string `toml:"db_name"`
}

var AppConfig Config
