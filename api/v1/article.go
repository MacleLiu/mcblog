package v1

import (
	"mcblog/models"
	"mcblog/utils/errno"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加文章
func AddArticle(ctx *gin.Context) {
	var art models.Article
	if err := ctx.ShouldBindJSON(&art); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": errno.ERROR,
			"msg":    err.Error(),
		})
		return
	}

	id, err := models.CreateArticle(&art)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   id, // 添加成功返回id, 用于添加标签
		"msg":    errno.GetMsg(err),
	})
}

// 查询单个文章
func GetArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	art, err := models.GetArticle(id)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   art,
		"msg":    errno.GetMsg(err),
	})
}

// 查询文章列表
func GetArticles(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	title := ctx.Query("title")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	arts, total, err := models.GetArticles(pageSize, pageNum, title)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   arts,
		"total":  total,
		"msg":    errno.GetMsg(err),
	})
}

// 查询精选文章列表
func GetWinnowArticles(ctx *gin.Context) {
	arts, err := models.GetWinnowArticles()
	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   arts,
		"msg":    errno.GetMsg(err),
	})

}

// 查询全部文章的基础信息
func GetAllArtInfo(ctx *gin.Context) {
	arts, err := models.GetAllArtInfo()
	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   arts,
		"msg":    errno.GetMsg(err),
	})
}

// 查询文章总数
func GetArticleCount(ctx *gin.Context) {
	c, err := models.GetArticCuont()
	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   c,
		"msg":    errno.GetMsg(err),
	})
}

// 查询分类下的文章
func GetCateArticles(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	cid, _ := strconv.Atoi(ctx.Param("id"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	arts, total, err := models.GetCateArticles(cid, pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   arts,
		"total":  total,
		"msg":    errno.GetMsg(err),
	})
}

// 编辑文章
func EditArticle(ctx *gin.Context) {
	var art models.Article
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindJSON(&art); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": errno.ERROR,
			"msg":    err.Error(),
		})
		return
	}

	err := models.CheckArticleExist(id)
	if err == nil {
		err = models.EditArticle(id, art)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})

}

// 删除文章
func DeleteArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := models.CheckArticleExist(id)
	if err == nil {
		err = models.DeleteArticle(id)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})
}

// 设置精选文章
func SetWinnow(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := models.SetWinnow(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})
}
