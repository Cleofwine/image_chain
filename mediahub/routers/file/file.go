package file

import (
	"mediahub/controller/file"

	"github.com/gin-gonic/gin"
)

func InitFileRouter(api *gin.RouterGroup, fileController *file.File) {
	v1 := api.Group("/v1")
	fileGroup := v1.Group("/file")
	fileGroup.POST("/upload", fileController.Upload)
}
