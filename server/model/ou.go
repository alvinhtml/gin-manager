package model

import (
	"github.com/alvinhtml/gin-manager/server/global"
)

type Ou struct {
	global.MODEL
	Name        string `json:"name" gorm:"comment:组织单位名称;type:varchar(255);size:255;"`
	Description string `json:"description" gorm:"comment:描述;type:varchar(2000);size:2000;"`
	Pid         uint   `json:"pid" gorm:"comment:上级组织单位;type:int(64);size:64"`
}
