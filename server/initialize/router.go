package initialize

import (
	_ "gin-manager/docs"
	"gin-manager/global"
	"gin-manager/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()

	Router.StaticFS(global.CONFIG.Local.Path, http.Dir(global.CONFIG.Local.Path)) // 为用户头像和文件提供静态地址

	// Router.Use(middleware.LoadTls())  // 打开就能玩https了

	global.LOG.Info("use middleware logger")

	// 跨域
	// Router.Use(middleware.Cors())
	// global.LOG.Info("use middleware cors")

	// swagger 文档
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Info("register swagger handler")

	ApiGroup := Router.Group("/api")
	router.InitUserRouter(ApiGroup) // 注册用户路由
	router.InitOuRouter(ApiGroup)   // 注册部门路由

	global.LOG.Info("router register success")
	return Router
}
