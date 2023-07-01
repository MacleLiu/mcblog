package v1

import (
	"mcblog/upload"
	"mcblog/utils/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpLoad(ctx *gin.Context) {
	file, fileHeader, _ := ctx.Request.FormFile("file")

	fileSize := fileHeader.Size

	url, err := upload.UpLoadFile(file, fileSize)

	ctx.JSON(http.StatusOK, gin.H{
		"status": errno.GetCode(err),
		"msg":    errno.GetMsg(err),
		"url":    url,
	})
}
