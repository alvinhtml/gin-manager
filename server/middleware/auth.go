package middleware

import (
	"errors"
	"strings"

	"github.com/alvinhtml/gin-manager/server/global/response"
	"github.com/alvinhtml/gin-manager/server/service"
	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("X-Token")

		if token == "" {
			header := c.GetHeader("Authorization")
			headerList := strings.Split(header, " ")

			if len(headerList) != 2 {
				err := errors.New("无法解析 Authorization 字段")

				response.BadRequest(err, c)
				c.Abort()
				return
			}

			t := headerList[0]
			if t != "Bearer" {
				err := errors.New("Authorization 类型错误, 当前只支持 Bearer")

				response.BadRequest(err, c)
				c.Abort()
				return
			}

			token = headerList[1]
		}

		if err := service.VerifyToken(token); err != nil {
			err := errors.New("401 Unauthorized")
			response.Unauthorized(err, c)
			c.Abort()
			return
		}

		c.Next()
	}
}
