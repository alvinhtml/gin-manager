package initialize

import (
	"net/http"

	_ "github.com/alvinhtml/gin-manager/server/docs"
	"github.com/alvinhtml/gin-manager/server/global"
	"github.com/alvinhtml/gin-manager/server/middleware"
	"github.com/alvinhtml/gin-manager/server/router"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// 公共路由，无需 JWT 验证
	PublicGroup := Router.Group("/api")
	router.InitPublicRouter(PublicGroup)

	ApiGroup := Router.Group("/api")
	ApiGroup.Use(middleware.AuthJWT())
	router.InitUserRouter(ApiGroup) // 注册用户路由
	router.InitOuRouter(ApiGroup)   // 注册组织单位路由

	global.LOG.Info("router register success")
	return Router
}
