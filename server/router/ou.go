package router

import (
	v1 "github.com/alvinhtml/gin-manager/server/api/v1"

	"github.com/gin-gonic/gin"
)

// InitUserRouter
func InitOuRouter(Router *gin.RouterGroup) {
	OuRouter := Router.Group("ous")

	{
		OuRouter.GET("", v1.GetOus) // 分页获取部门列表
	}
}
