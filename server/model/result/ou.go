package result

import (
	"gin-manager/model"
)

type Ou struct {
	model.Ou
	ParentName string `json:"parent_name" from:"parent_name"`
}
