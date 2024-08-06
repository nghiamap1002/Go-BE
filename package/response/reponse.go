package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `mapstructure:"code"`
	Message string      `mapstructure:"message"`
	Data    interface{} `mapstructure:"data"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusBadRequest, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
