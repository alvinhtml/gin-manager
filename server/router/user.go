package router

import (
	v1 "gin-manager/api/v1"

	"github.com/gin-gonic/gin"
)

// InitUserRouter
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("users")

	{
		UserRouter.GET("", v1.GetUserList) // 分页获取用户列表
	}
}
