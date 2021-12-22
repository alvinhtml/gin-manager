package router

import (
	v1 "github.com/alvinhtml/gin-manager/server/api/v1"

	"github.com/gin-gonic/gin"
)

// InitUserRouter
func InitPublicRouter(Router *gin.RouterGroup) {
	Router.POST("/login", v1.Login)
	Router.POST("/register", v1.Register)
}
