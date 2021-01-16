package core

import (
	"fmt"
	"time"

	"gin-manager/global"
	"gin-manager/initialize"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 gin-manager
	当前版本: v0.1.0
	默认自动化文档地址: http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址: http://127.0.0.1:8080
`, address)
	global.LOG.Error(s.ListenAndServe().Error())
}
