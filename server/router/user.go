package router

import (
	v1 "github.com/alvinhtml/gin-manager/server/api/v1"

	"github.com/gin-gonic/gin"
)

// https://github.com/swaggo/swag

// InitUserRouter
func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("users")

	{
		UserRouter.GET("", v1.GetUsers)         // 分页获取用户列表
		UserRouter.GET(":id", v1.GetUser)       // 获取用户
		UserRouter.POST("", v1.CreateUser)      // 创建用户
		UserRouter.PUT(":id", v1.UpdateUser)    // 更新用户
		UserRouter.DELETE(":id", v1.DeleteUser) // 删除用户
	}
}
