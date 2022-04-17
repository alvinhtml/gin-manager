package result

import (
	"github.com/alvinhtml/gin-manager/server/model"
)

type UserProfile struct {
	model.UserJoinOu
	IsLogin bool `json:"is_login" from:"is_login"`
}
