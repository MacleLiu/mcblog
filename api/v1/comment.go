package v1

import (
	"mcblog/models"
	"mcblog/utils/errno"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加评论
func AddComment(ctx *gin.Context) {
	var comt models.Comment
	if err := ctx.ShouldBindJSON(&comt); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": errno.ERROR,
			"msg":    err.Error(),
		})
		return
	}

	id, err := models.CreateComment(&comt)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   id, // 添加成功返回评论id
		"msg":    errno.GetMsg(err),
	})
}

// 查询文章的评论
func GetComment(ctx *gin.Context) {
	type comment struct {
		models.Comment                  // 一级评论
		Children       []models.Comment `json:"children"` // 二级评论数组
	}
	var index int                  // 一级评论在comments中的下标
	idToIndex := make(map[int]int) // 将一级评论的id映射到其在评论数组的下标
	comments := make([]comment, 0)
	id, _ := strconv.Atoi(ctx.Param("id"))
	comts, total, err := models.GetCommentByArt(id)

	// 处理评论数据，将二级评论放入一级评论的Children
	// comts已按照id升序排序，根据评论的逻辑可知，一级评论的id一定小于二级评论
	// 所以升序排序后进行遍历，一定可以先访问到二级评论的一级评论
	// 而且按照id升序排序后，同时也隐式地按照了时间升序排序，因为id越大的评论其发布评论的时间也更晚，所以评论在时间上也是有序的。
	for _, comt := range comts {
		if comt.Toid == 0 {
			idToIndex[int(comt.ID)] = index
			index++
			comments = append(comments, comment{
				comt,
				[]models.Comment{},
			})
		} else {
			comments[idToIndex[comt.Toid]].Children = append(comments[idToIndex[comt.Toid]].Children, comt)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"data":   comments,
		"total":  total,
		"msg":    errno.GetMsg(err),
	})
}
