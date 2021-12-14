package global

import (
	"go.uber.org/zap"

	"github.com/alvinhtml/gin-manager/server/config"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	// DB 数据库句柄
	DB *gorm.DB

	// REDIS REDIS句柄
	REDIS *redis.Client

	// CONFIG config 配置
	CONFIG config.Server

	// VP  Viper句柄
	VP *viper.Viper

	// LOG  日志句柄
	LOG *zap.Logger
)
