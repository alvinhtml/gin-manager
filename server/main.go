package main

import (
	"github.com/alvinhtml/gin-manager/server/core"
	"github.com/alvinhtml/gin-manager/server/global"
	"github.com/alvinhtml/gin-manager/server/initialize"
)

// @title 			Swagger Example API
// @version 		0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in 					header
// @name 				x-token
// @BasePath 		/api
func main() {
	global.VP = core.Viper()          // 初始化Viper
	global.LOG = core.Zap()           // 初始化zap日志库
	global.DB = initialize.Gorm()     // gorm连接数据库
	initialize.MysqlTables(global.DB) // 初始化表

	// 程序结束前关闭数据库链接
	db, _ := global.DB.DB()
	defer db.Close()

	core.RunServer()
}
