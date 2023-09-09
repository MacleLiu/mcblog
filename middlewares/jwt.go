package middlewares

import (
	"fmt"
	"mcblog/config"
	"mcblog/utils/errno"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte(config.AppConfig.JwtSecret)

type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

// 生成 token
func GenerateToken(username string) (string, error) {
	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "mcblog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtSecret)
	return tokenStr, err
}

// 解析token
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errno.New(errno.ERROR_TOKEN_WRONG, nil)
	}
}

// JWT鉴权
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取请求头携带的Token
		tokenStr := ctx.Request.Header.Get("Authorization")
		if tokenStr == "" {
			err := errno.New(errno.ERROR_TOKEN_NOT_EXIST, nil)
			ctx.JSON(http.StatusOK, gin.H{
				"statsu": errno.GetCode(err),
				"msg":    errno.GetMsg(err),
			})
			ctx.Abort()
			return
		}

		// 检查token格式
		checkType := strings.SplitN(tokenStr, " ", 2)
		if len(checkType) != 2 || checkType[0] != "Bearer" {
			err := errno.New(errno.ERROR_TOKEN_TYPE_WRONG, nil)
			ctx.JSON(http.StatusOK, gin.H{
				"statsu": errno.GetCode(err),
				"msg":    errno.GetMsg(err),
			})
			ctx.Abort()
			return
		}

		// 解析客户端发送的 token
		_, err := ParseToken(checkType[1])
		if err != nil {

			fmt.Println(err)

			err = errno.New(errno.ERROR_TOKEN_WRONG, err)
			ctx.JSON(http.StatusOK, gin.H{
				"status": errno.GetCode(err),
				"msg":    errno.GetMsg(err),
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
