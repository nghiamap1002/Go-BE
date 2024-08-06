package user

import "github.com/gin-gonic/gin"

type ProductRouter struct {
}

func (pr *ProductRouter) InitProductRouter(r *gin.RouterGroup) {
	productRouterPublic := r.Group("/product")
	{
		productRouterPublic.GET("/search")
		productRouterPublic.GET("/detail/:id")
	}
}
