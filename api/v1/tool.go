package v1

import (
	"mcblog/models"
	"mcblog/utils/errno"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加工具
func AddTool(ctx *gin.Context) {
	var tool models.Tool
	if err := ctx.ShouldBindJSON(&tool); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": errno.ERROR,
			"msg":    err.Error(),
		})
		return
	}
	err := models.CreateTool(&tool)
	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})
}

// 分页查询工具列表
func GetTools(ctx *gin.Context) {
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

	tools, total, err := models.GetTools(pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   tools,
		"total":  total,
		"msg":    errno.GetMsg(err),
	})
}
