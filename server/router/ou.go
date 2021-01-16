package router

import (
	v1 "gin-manager/api/v1"

	"github.com/gin-gonic/gin"
)

// InitUserRouter
func InitOuRouter(Router *gin.RouterGroup) {
	// UserRouter := Router.Group("user")

	Router.GET("ous", v1.GetOuList) // 分页获取部门列表
}
