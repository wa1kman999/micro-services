package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wa1kman999/helloWorld-client/controller/user"
)

func InitRouter(router *gin.Engine) {
	userGroup := router.Group("/user")

	// 调用
	user.Router(userGroup)
}
