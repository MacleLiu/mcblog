package v1

import (
	"mcblog/models"
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
	var cate models.Category
	if err := ctx.ShouldBindJSON(&cate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": errno.ERROR,
			"msg":    err.Error(),
		})
		return
	}

	// 检查分类是否已存在
	err := models.CheckCategoryName(cate.Name)
	if err == nil {
		err = errno.New(errno.ERROR_CATE_USED, err)
	} else if errno.GetCode(err) == errno.ERROR_CATE_NOT_EXIST {
		err = models.CreateCategory(&cate)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})
}

// 查询单个分类信息
func GetCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	cate, err := models.GetCategory(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   cate,
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

	cates, total, err := models.GetCategories(pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   cates,
		"total":  total,
		"msg":    errno.GetMsg(err),
	})
}

// 查询分类统计信息（分类下文章数量）
func GetCateStat(ctx *gin.Context) {
	cateStat, err := models.GetCateStat()
	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   cateStat,
		"msg":    errno.GetMsg(err),
	})
}

// 编辑分类
func EditCategory(ctx *gin.Context) {
	var cate models.Category
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindJSON(&cate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": errno.ERROR,
			"msg":    err.Error(),
		})
		return
	}

	// 检查目标分类是否存在
	if err := models.CheckCategory(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errno.GetCode(err),
			"msg":    errno.GetMsg(err),
		})
		return
	}

	// 检查分类名是否已使用
	err := models.CheckCategoryName(cate.Name)
	if err == nil {
		err = errno.New(errno.ERROR_CATE_USED, err)
	} else if errno.GetCode(err) == errno.ERROR_CATE_NOT_EXIST {
		err = models.EditCategory(id, cate)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})

}

// 删除分类
func DeleteCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// 检查目标分类是否存在
	if err := models.CheckCategory(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errno.GetCode(err),
			"msg":    errno.GetMsg(err),
		})

		return
	}

	// 检查分类是否正在被使用
	if err := models.CheckCateUsed(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errno.GetCode(err),
			"msg":    errno.GetMsg(err),
		})
		return
	}

	err := models.DeleteCategory(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})
}
