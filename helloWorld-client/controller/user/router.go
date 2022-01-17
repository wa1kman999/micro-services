package user

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	// 用户登录
	router.POST("/login", Login)
}
