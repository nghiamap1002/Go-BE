package routers

import (
	"fmt"
	"personal/ShopDev/Go-BE/internal/controller"
	"personal/ShopDev/Go-BE/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before AA")
		c.Next()
		fmt.Println("After AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before BB")
		c.Next()
		fmt.Println("After BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before CC")
	c.Next()
	fmt.Println("After CC")
}

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.AuthNiddleware(), AA(), BB(), CC)

	v1 := r.Group("v1")
	{
		v1.GET("/ping", controller.NewPongCotroller().Pong)
		v1.GET("/ping/:name", controller.NewUserController().GetUserById)
	}

	return r
}
