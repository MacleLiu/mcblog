package v1

import (
	"mcblog/models"
	"mcblog/utils/errno"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加心愿
func AddWish(ctx *gin.Context) {
	var wish models.Wish
	if err := ctx.ShouldBindJSON(&wish); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": errno.ERROR,
			"msg":    err.Error(),
		})
		return
	}
	err := models.CreateWish(&wish)
	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})
}

// 查询单个心愿信息
func GetWish(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	wish, err := models.GetWish(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   wish,
		"msg":    errno.GetMsg(err),
	})
}

// 分页查询心愿列表
func GetWishes(ctx *gin.Context) {
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

	wishs, total, err := models.GetWishes(pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   wishs,
		"total":  total,
		"msg":    errno.GetMsg(err),
	})
}

// 编辑心愿
func EditWish(ctx *gin.Context) {
	var tool models.Wish
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindJSON(&tool); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": errno.ERROR,
			"msg":    err.Error(),
		})
		return
	}

	err := models.EditWish(id, tool)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})

}

// 删除心愿
func DeleteWish(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// 检查目标心愿是否存在
	// :todo

	err := models.DeleteWish(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})
}
