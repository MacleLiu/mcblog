package v1

import (
	"mcblog/modles"
	"mcblog/utils/errno"
	"mcblog/utils/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查询用户是否存在
func UserExist(ctx *gin.Context) {

}

// 添加用户
func AddUser(ctx *gin.Context) {
	var user modles.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": errno.ERROR,
			"msg":    err.Error(),
		})
		return
	}

	//字段合法性校验
	if msg, err := validator.Validator(&user); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errno.GetCode(err),
			"msg":    msg,
		})

		return
	}

	err := modles.CheckUserName(user.Username)
	if err == nil {
		err = modles.CreateUser(&user)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})
}

//查询单个用户

// 查询用户列表
func GetUsers(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	users, err := modles.GetUsers(pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   users,
		"msg":    errno.GetMsg(err),
	})
}

// 编辑用户
func EditUser(ctx *gin.Context) {
	var user modles.User
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": errno.ERROR,
			"msg":    err.Error(),
		})
		return
	}

	//字段合法性校验
	if msg, err := validator.VarValidator(user.Username, "required,min=4,max=12"); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errno.GetCode(err),
			"msg":    "用户名" + msg,
		})
		return
	}
	if msg, err := validator.VarValidator(user.Role, "required,gte=2"); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errno.GetCode(err),
			"msg":    "角色码" + msg,
		})
		return
	}

	//检查用户名是否已存在
	err := modles.CheckUserName(user.Username)
	if err == nil {
		err = modles.EditUser(id, user)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})

}

// 删除用户
func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := modles.DeleteUser(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})
}
