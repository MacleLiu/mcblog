package v1

import (
	"mcblog/models"
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
	var user models.User
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

	//检查用户名是否已使用
	err := models.CheckUserName(user.Username)
	if errno.GetCode(err) == errno.ERROR_USER_NOT_EXIST {
		err = models.CreateUser(&user)
	} else if err == nil {
		err = errno.New(errno.ERROR_USERNAME_USED, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})
}

// 查询单个用户
func GetUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user, err := models.GetUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   user,
		"msg":    errno.GetMsg(err),
	})
}

// 查询用户列表
func GetUsers(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	username := ctx.Query("username")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	users, total, err := models.GetUsers(pageSize, pageNum, username)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   users,
		"total":  total,
		"msg":    errno.GetMsg(err),
	})
}

// 编辑用户
func EditUser(ctx *gin.Context) {
	var user models.User
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
	if msg, err := validator.VarValidator(user.Role, "required,oneof=1 2"); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errno.GetCode(err),
			"msg":    "角色码" + msg,
		})
		return
	}

	// 检查目标用户是否存在
	if err := models.CheckUser(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errno.GetCode(err),
			"msg":    errno.GetMsg(err),
		})

		return
	}

	// 检查新用户名是否被其他用户使用
	err := models.CheckUpUser(id, user.Username)
	if err == nil || errno.GetCode(err) == errno.ERROR_USER_NOT_EXIST {
		err = models.EditUser(id, user)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})

}

// 删除用户
func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	//检查目标用户是否存在
	if err := models.CheckUser(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errno.GetCode(err),
			"msg":    errno.GetMsg(err),
		})

		return
	}

	err := models.DeleteUser(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})
}
