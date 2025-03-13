package home

import (
	"mediahub/controller/home"

	"github.com/gin-gonic/gin"
)

func InitHomeRouter(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	v1.GET("/home", home.Home)
}
