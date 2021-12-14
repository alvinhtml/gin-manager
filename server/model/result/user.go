package result

import (
	"github.com/alvinhtml/gin-manager/server/model"
)

type User struct {
	model.User
	OuName string `json:"ou_name" from:"ou_name"`
}
