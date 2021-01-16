package result

import (
	"gin-manager/model"
)

type User struct {
	model.User
	OuName string `json:"ou_name" from:"ou_name"`
}
