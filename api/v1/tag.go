package v1

import (
	"mcblog/models"
	"mcblog/utils/errno"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加标签
func AddTag(ctx *gin.Context) {
	var tag models.Tag
	if err := ctx.ShouldBindJSON(&tag); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": errno.ERROR,
			"msg":    err.Error(),
		})
		return
	}

	err := models.CheckTagName(tag.Name)
	if err == nil {
		err = errno.New(errno.ERROR_TAG_USED, err)
	} else if errno.GetCode(err) == errno.ERROR_TAG_NOT_EXIST {
		err = models.CreateTag(&tag)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})
}

// 查询标签列表
func GetTags(ctx *gin.Context) {
	tags, total, err := models.GetTags()

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   tags,
		"total":  total,
		"msg":    errno.GetMsg(err),
	})
}

// 分页查询标签下的文章
func GetTagArticles(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	tid, _ := strconv.Atoi(ctx.Param("id"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	tags, total, err := models.GetTagArticles(tid, pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   tags,
		"total":  total,
		"msg":    errno.GetMsg(err),
	})
}

// 查询一个标签
func GetTag(ctx *gin.Context) {
	tid, _ := strconv.Atoi(ctx.Param("id"))

	art, err := models.GetTag(tid)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   art,
		"msg":    errno.GetMsg(err),
	})
}

// 查询文章标签
func GetArticleTags(ctx *gin.Context) {
	aid, _ := strconv.Atoi(ctx.Param("id"))

	art, err := models.GetArticleTags(aid)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   art,
		"msg":    errno.GetMsg(err),
	})
}

// 更新文章标签
func UpdateArticleTags(ctx *gin.Context) {
	var tags []models.Tag
	aid, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindJSON(&tags); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": errno.ERROR,
			"msg":    err.Error(),
		})
		return
	}
	err := models.UpdateArticleTags(aid, tags)
	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
	})
}
