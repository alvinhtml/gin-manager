package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
const (
    StatusContinue           = 100
    StatusSwitchingProtocols = 101
    StatusOK                   = 200
    StatusCreated              = 201
    StatusAccepted             = 202
    StatusNonAuthoritativeInfo = 203
    StatusNoContent            = 204
    StatusResetContent         = 205
    StatusPartialContent       = 206
    StatusMultipleChoices   = 300
    StatusMovedPermanently  = 301
    StatusFound             = 302
    StatusSeeOther          = 303
    StatusNotModified       = 304
    StatusUseProxy          = 305
    StatusTemporaryRedirect = 307
    StatusBadRequest                   = 400
    StatusUnauthorized                 = 401
    StatusPaymentRequired              = 402
    StatusForbidden                    = 403
    StatusNotFound                     = 404
    StatusMethodNotAllowed             = 405
    StatusNotAcceptable                = 406
    StatusProxyAuthRequired            = 407
    StatusRequestTimeout               = 408
    StatusConflict                     = 409
    StatusGone                         = 410
    StatusLengthRequired               = 411
    StatusPreconditionFailed           = 412
    StatusRequestEntityTooLarge        = 413
    StatusRequestURITooLong            = 414
    StatusUnsupportedMediaType         = 415
    StatusRequestedRangeNotSatisfiable = 416
    StatusExpectationFailed            = 417
    StatusTeapot                       = 418
    StatusInternalServerError     = 500
    StatusNotImplemented          = 501
    StatusBadGateway              = 502
    StatusServiceUnavailable      = 503
    StatusGatewayTimeout          = 504
    StatusHTTPVersionNotSupported = 505
)
*/

// Message 响应体消息
type Message struct {
	Message string `json:"message" form:"message"`
}

// Success 成功，返回JSON数据
func Success(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, data)
}

// NotFound 404
func NotFound(data interface{}, c *gin.Context) {
	c.JSON(http.StatusNotFound, Message{
		Message: "not find",
	})
}

// Fail 失败
func Fail(err error, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Message{
		Message: err.Error(),
	})
}
