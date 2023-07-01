package config

type Config struct {
	Server    Server `toml:"server"`
	Mysql     Mysql  `toml:"database"`
	JwtSecret string `toml:"jwtSecret"`
	QiNiu     QiNiu  `toml:"qiniu"`
}

type Server struct {
	AppModel string `toml:"app_model"`
	AppPort  string `toml:"app_port"`
}

type Mysql struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DbName   string `toml:"db_name"`
}

type QiNiu struct {
	AccessKey   string `toml:"access_key"`
	SecretKey   string `toml:"secret_key"`
	Bucket      string `toml:"bucket"`
	QiNiuServer string `toml:"qiniu_server"`
}

var AppConfig Config
