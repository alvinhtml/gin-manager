package router

import (
	v1 "github.com/alvinhtml/gin-manager/server/api/v1"

	"github.com/gin-gonic/gin"
)

// InitUserRouter
func InitOuRouter(Router *gin.RouterGroup) {
	OuRouter := Router.Group("ous")

	{
		OuRouter.GET("", v1.GetOus)          // 获取组织单位列表
		OuRouter.GET("/:id", v1.GetOu)       // 获取组织单位详情
		OuRouter.POST("", v1.CreateOu)       // 创建组织单位
		OuRouter.PUT("/:id", v1.UpdateOu)    // 更新组织单位
		OuRouter.DELETE("/:id", v1.DeleteOu) // 删除组织单位
	}
}
