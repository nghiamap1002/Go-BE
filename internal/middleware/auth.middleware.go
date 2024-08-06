package middleware

import (
	"fmt"
	"personal/ShopDev/Go-BE/package/response"

	"github.com/gin-gonic/gin"
)

func AuthNiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token != "valid-token" {
			fmt.Println("valid-token")
			response.ErrorResponse(ctx, response.ErrorInvalidToken, "")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
