package result

import "time"

type PageResult struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Size  int         `json:"pageSize"`
}

type Token struct {
	Token  string `json:"token"`
	expire time.Time
}
