package router

import (
	v1 "github.com/alvinhtml/gin-manager/server/api/v1"

	"github.com/gin-gonic/gin"
)

// InitUserRouter
func InitOuRouter(Router *gin.RouterGroup) {
	OuRouter := Router.Group("ous")

	{
		OuRouter.GET("", v1.GetOus)          // 获取部门列表
		OuRouter.GET("/:id", v1.GetOu)       // 获取部门详情
		OuRouter.POST("", v1.CreateOu)       // 创建部门
		OuRouter.PUT("/:id", v1.UpdateOu)    // 更新部门
		OuRouter.DELETE("/:id", v1.DeleteOu) // 删除部门
	}
}
