package model

import (
	"github.com/alvinhtml/gin-manager/server/global"
)

type Jwt struct {
	global.MODEL
	Jwt string `json:"name" gorm:"comment:jwt;type:text;"`
}
