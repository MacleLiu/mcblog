package v1

import (
	"mcblog/middlewares"
	"mcblog/models"
	"mcblog/utils/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var user models.User
	var token string
	var err error
	if err = ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": errno.ERROR,
			"msg":    err.Error(),
		})
		return
	}

	if err = models.LoginVerify(user.Username, user.Password); err == nil {
		token, err = middlewares.GenerateToken(user.Username)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
		"token":  token,
	})

}
