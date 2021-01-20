package v1

import (
	"gin-manager/global/response"
	"gin-manager/model/request"
	"gin-manager/model/result"
	"gin-manager/service"

	"github.com/gin-gonic/gin"
)

// @Tags apiUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo

	err, list, total := service.GetUserInfoList(pageInfo)
	if err != nil {
		response.Fail(err, c)
	} else {
		response.Success(result.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
