package v1

import (
	"fmt"

	"gin-manager/global/response"
	"gin-manager/model/request"
	"gin-manager/model/result"
	"gin-manager/service"

	"github.com/gin-gonic/gin"
)

// @Tags apiOu
// @Summary 分页获取部门列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取部门列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func GetOuList(c *gin.Context) {
	var pageInfo request.PageInfo

	err, list, total := service.GetOuList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(result.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
