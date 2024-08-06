package user

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	userRouterPublic := r.Group("/user")
	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/otp")
	}

	userRouterPrivate := r.Group("/user")
	{
		userRouterPrivate.POST("/register")
		userRouterPrivate.POST("/otp")
	}
}
