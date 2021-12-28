package v1

import (
	"fmt"

	"github.com/alvinhtml/gin-manager/server/global/response"
	"github.com/alvinhtml/gin-manager/server/model/request"
	"github.com/alvinhtml/gin-manager/server/service"

	"github.com/alvinhtml/gin-manager/server/model"

	"github.com/gin-gonic/gin"
)

// @Tags 			apiPublic
// @Summary 	用户登录
// @accept 		application/json
// @Produce 	application/json
// @Param 		data body request.LoginForm true "用户名和密码"
// @Success 	200 {object} model.Jwt "jwt token"
// @Router 		/login [post]
func Login(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)

	err := service.Login(user)
	fmt.Println(user)

	token, err := service.CreateToken(user.Name)

	if err != nil {
		response.Fail(err, c)
	} else {

		response.Success(token, c)
	}
}

// @Tags 			apiPublic
// @Summary 	用户注册
// @accept 		application/json
// @Produce 	application/json
// @Param 		data body model.User true "用户信息"
// @Success 	200 {object} model.User "用户详情"
// @Router 		/register [post]
func Register(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)

	err, user := service.CreateUser(user)
	if err != nil {
		response.Fail(err, c)
	} else {
		response.Success(user, c)
	}
}

func Test(c *gin.Context) {

	var pageInfo request.PageQuery
	err := c.ShouldBindQuery(&pageInfo)
	fmt.Println(pageInfo)

	if err != nil {
		response.Fail(err, c)
	} else {
		response.Success(pageInfo, c)
	}
}
