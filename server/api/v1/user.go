package v1

import (
	"github.com/alvinhtml/gin-manager/server/global/response"
	"github.com/alvinhtml/gin-manager/server/model/request"
	"github.com/alvinhtml/gin-manager/server/model/result"
	"github.com/alvinhtml/gin-manager/server/service"

	"github.com/alvinhtml/gin-manager/server/model"

	"github.com/gin-gonic/gin"
)

// @Tags 			apiUser
// @Summary 	获取用户列表
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce 	application/json
// @Param 		page query string false "当前页码"
// @Param 		size query string false "每页显示条数"
// @success 	200 {object} result.PageResult{list=[]model.Ou} "用户列表"
// @Router 		/users [get]
func GetUsers(c *gin.Context) {
	var pageInfo request.PageInfo
	c.ShouldBindJSON(&pageInfo)

	err, list, total := service.GetUsers(pageInfo)
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

// @Tags 			apiUser
// @Summary 	获取用户详情
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce 	application/json
// @Param 		id path string true "id"
// @Success 	200 {object} model.User "用户详情"
// @Router 		/users/{id} [get]
func GetUser(c *gin.Context) {
	var byId request.ById
	c.ShouldBindUri(&byId)

	err, user := service.GetUser(byId.Id)

	if err != nil {
		response.Fail(err, c)
	} else {
		response.Success(user, c)
	}
}

// @Tags 			apiUser
// @Summary 	更新用户
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce 	application/json
// @Param 		data body model.User true "要修改的用户信息"
// @Success 	200 {object} model.User "用户详情"
// @Router 		/users/{id} [post]
func UpdateUser(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)

	var byId request.ById
	c.ShouldBindJSON(&byId)

	user.Id = byId.Id

	err, _ := service.UpdateUser(user)
	if err != nil {
		response.Fail(err, c)
	} else {
		response.Success(nil, c)
	}
}

//	@Tags 			apiUser
//	@Summary 		删除用户
//	@Security 	ApiKeyAuth
//	@accept 		application/json
//	@Produce 		application/json
//	@Param 			id path string true "id"
//	@Success 		200 {object} model.User "用户详情"
//	@Router 		/users/{id} [delete]
func DeleteUser(c *gin.Context) {
	var byId request.ById
	c.ShouldBindUri(&byId)

	err := service.DeleteUser(byId.Id)
	if err != nil {
		response.Fail(err, c)
	} else {
		response.Success(nil, c)
	}
}

// @Tags 			apiUser
// @Summary 	创建用户
// @Security 	ApiKeyAuth
// @accept 		application/json
// @Produce 	application/json
// @Param 		data body model.User true "要创建的用户信息"
// @Success 	200 {object} model.User "用户详情"
// @Router 		/users [post]
func CreateUser(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)

	err, _ := service.CreateUser(user)
	if err != nil {
		response.Fail(err, c)
	} else {
		response.Success(nil, c)
	}
}
