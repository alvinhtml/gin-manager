package model

import (
	"github.com/alvinhtml/gin-manager/server/global"
)

type User struct {
	global.MODEL
	Name     string `json:"name" gorm:"comment:用户登录名;type:varchar(255);size:255;"`
	Password string `json:"-"  gorm:"comment:用户登录密码;type:char(64);size:64;"`
	Email    string `json:"email" gorm:"comment:邮箱;type:varchar(100);size:100;"`
	Profile  string `json:"headerImg" gorm:"comment:用户头像;type:varchar(255);size:255;"`
	Status   int    `json:"status" gorm:"comment:状态;type:int(1);size:1;"`
	OuRefer  uint   `json:"ou_id"`
}

type UserWithOu struct {
	User   `json:"instance"`
	OuName string `json:"ou_name"`
}
