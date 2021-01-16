package router

import (
	v1 "gin-manager/api/v1"

	"github.com/gin-gonic/gin"
)

// InitUserRouter
func InitUserRouter(Router *gin.RouterGroup) {
	// UserRouter := Router.Group("user")

	Router.GET("users", v1.GetUserList) // 分页获取用户列表
}
