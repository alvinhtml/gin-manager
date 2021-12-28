package result

import (
	"time"

	"github.com/alvinhtml/gin-manager/server/model/request"
)

type PageResult struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
	request.PageQuery
}

type Token struct {
	Token  string `json:"token"`
	expire time.Time
}
