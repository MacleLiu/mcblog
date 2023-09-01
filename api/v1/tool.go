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

// 查询单个工具信息
func GetTool(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	tool, err := models.GetTool(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   tool,
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

// 编辑工具
func EditTool(ctx *gin.Context) {
	var tool models.Tool
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindJSON(&tool); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": errno.ERROR,
			"msg":    err.Error(),
		})
		return
	}

	err := models.EditTool(id, tool)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})

}

// 删除工具
func DeleteTool(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// 检查目标工具是否存在
	// :todo

	err := models.DeleteTool(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})
}
