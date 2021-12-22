package model

import (
	"time"

	"github.com/alvinhtml/gin-manager/server/global"
)

type Jwt struct {
	global.MODEL
	Token     string    `json:"token" gorm:"comment:token;type:text;"`
	ExpiresAt time.Time `json:"ExpiresAt" gorm:"comment:ExpiresAt;type:timestamp;"`
}
