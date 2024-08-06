package manage

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	// userRouterPublic := r.Group("/user")
	// {
	// 	userRouterPublic.POST("/register")
	// 	userRouterPublic.POST("/otp")
	// }

	userRouterPrivate := r.Group("/admin/user")
	{
		userRouterPrivate.POST("/active_user")
	}
}
