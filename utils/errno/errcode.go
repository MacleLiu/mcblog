package errno

const (
	SUCCESS = 200
	ERROR   = 500

	//code=1000...用户模块错误
	ERROR_USERNAME_USED      = 1001
	ERROR_PASSWORD_WRONG     = 1002
	ERROR_USER_NOT_EXIXT     = 1003
	ERROR_TOKEN_NOT_EXIST    = 1004
	ERROR_TOKEN_TIMEOUT      = 1005
	ERROR_TOKEN_WRONG        = 1006
	ERROR_TOKEN_FORMAT_WRONG = 1007
	//code=2000... 文章模块错误

	//code=3000... 分类模块错误

)

var codemsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	ERROR_USERNAME_USED:      "用户名已存在",
	ERROR_PASSWORD_WRONG:     "密码错误",
	ERROR_USER_NOT_EXIXT:     "用户不存在",
	ERROR_TOKEN_NOT_EXIST:    "TOKEN不存在",
	ERROR_TOKEN_TIMEOUT:      "TOKEN已过期",
	ERROR_TOKEN_WRONG:        "TOKEN错误",
	ERROR_TOKEN_FORMAT_WRONG: "TOKEN格式错误",
}
