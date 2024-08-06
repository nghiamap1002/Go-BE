package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (pr *AdminRouter) InitAdminRouter(r *gin.RouterGroup) {
	adminRouterPublic := r.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	adminRouterPrivate := r.Group("/admin/user")
	{
		adminRouterPrivate.POST("/active_user")
	}
}
