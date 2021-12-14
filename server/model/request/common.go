package request

// Paging common input parameter structure
type PageInfo struct {
	Page int `json:"page" form:"page"`
	Size int `json:"size" form:"size"`
}

// Find by id structure
type ById struct {
	Id uint `json:"id" form:"id" uri:"id"`
}

type byIds struct {
	Ids []uint `json:"ids" form:"ids"`
}
