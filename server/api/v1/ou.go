package v1

import (
	"github.com/alvinhtml/gin-manager/server/global/response"
	"github.com/alvinhtml/gin-manager/server/model"
	"github.com/alvinhtml/gin-manager/server/model/request"
	"github.com/alvinhtml/gin-manager/server/model/result"
	"github.com/alvinhtml/gin-manager/server/service"
	"github.com/gin-gonic/gin"
)

// @Tags			apiOu
// @Summary 	获取部门列表
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param 		page query string false "当前页码"
// @Param 		size query string false "每页显示条数"
// @success 	200 {object} result.PageResult{list=[]model.Ou} "部门列表"
// @Router 		/ous [get]
func GetOus(c *gin.Context) {
	var pageInfo request.PageInfo
	c.ShouldBindJSON(&pageInfo)

	err, list, total := service.GetOus(pageInfo)
	if err != nil {
		response.Fail(err, c)
	} else {
		response.Success(result.PageResult{
			List:  list,
			Total: total,
			Page:  pageInfo.Page,
			Size:  pageInfo.Size,
		}, c)
	}
}

// @Tags			apiOu
// @Summary 	获取部门详情
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param 		id path string true "部门id"
// @success 	200 {object} model.Ou "desc"
// @Router 		/ous/{id} [get]
func GetOu(c *gin.Context) {
	var byId request.ById
	c.ShouldBindUri(&byId)

	err, ou := service.GetOu(byId.Id)
	if err != nil {
		response.Fail(err, c)
	} else {
		response.Success(ou, c)
	}
}

// @Tags			apiOu
// @Summary 	创建部门
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param 		body body model.Ou true "部门信息"
// @success 	200 {object} model.Ou "desc"
// @Router 		/ous [post]
func CreateOu(c *gin.Context) {
	var ou model.Ou
	c.ShouldBindJSON(&ou)

	err, ou := service.CreateOu(ou)
	if err != nil {
		response.Fail(err, c)
	} else {
		response.Success(ou, c)
	}
}

// @Tags			apiOu
// @Summary 	更新部门
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param 		id path string true "部门id"
// @Param 		body body model.Ou true "部门信息"
// @success 	200 {object} model.Ou "desc"
// @Router 		/ous/{id} [put]
func UpdateOu(c *gin.Context) {
	var byId request.ById
	c.ShouldBindUri(&byId)

	var ou model.Ou
	c.ShouldBindJSON(&ou)

	ou.Id = byId.Id

	err, ou := service.UpdateOu(ou)
	if err != nil {
		response.Fail(err, c)
	} else {
		response.Success(ou, c)
	}
}

// @Tags			apiOu
// @Summary 	删除部门
// @Security	ApiKeyAuth
// @accept		application/json
// @Produce		application/json
// @Param 		id path string true "部门id"
// @success 	200 {object} model.Ou "desc"
// @Router 		/ous/{id} [delete]
func DeleteOu(c *gin.Context) {
	var byId request.ById
	c.ShouldBindUri(&byId)

	err := service.DeleteOu(byId.Id)
	if err != nil {
		response.Fail(err, c)
	} else {
		response.Success(nil, c)
	}
}
