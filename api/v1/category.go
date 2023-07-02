package v1

import (
	"mcblog/modles"
	"mcblog/utils/errno"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查询分类是否存在
func CategoryExist(ctx *gin.Context) {

}

// 添加分类
func AddCategory(ctx *gin.Context) {
	var cate modles.Category
	if err := ctx.ShouldBindJSON(&cate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": errno.ERROR,
			"msg":    err.Error(),
		})
		return
	}

	err := modles.CheckCategoryName(cate.Name)
	if err == nil {
		err = errno.New(errno.ERROR_CATE_USED, err)
	} else if errno.GetCode(err) == errno.ERROR_CATE_NOT_EXIST {
		err = modles.CreateCategory(&cate)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})
}

// 查询分类列表
func GetCategories(ctx *gin.Context) {
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

	cates, total, err := modles.GetCategories(pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   cates,
		"total":  total,
		"msg":    errno.GetMsg(err),
	})
}

// 编辑分类
func EditCategory(ctx *gin.Context) {
	var cate modles.Category
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindJSON(&cate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": errno.ERROR,
			"msg":    err.Error(),
		})
		return
	}

	// 检查目标分类是否存在
	if err := modles.CheckCategory(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errno.GetCode(err),
			"msg":    errno.GetMsg(err),
		})

		return
	}

	err := modles.EditCategory(id, cate)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})

}

// 删除分类
func DeleteCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// 检查目标分类是否存在
	if err := modles.CheckCategory(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errno.GetCode(err),
			"msg":    errno.GetMsg(err),
		})

		return
	}

	err := modles.DeleteCategory(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})
}
