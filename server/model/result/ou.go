package result

import (
	"github.com/alvinhtml/gin-manager/server/model"
)

type Ou struct {
	model.Ou
	ParentName string `json:"parent_name" from:"parent_name"`
}
