package initialize

import (
	"personal/ShopDev/Go-BE/global"
	"personal/ShopDev/Go-BE/internal/routers"

	"github.com/gin-gonic/gin"
)

// đi ngược next with nodejs

// func AA() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before AA")
// 		c.Next()
// 		fmt.Println("After AA")
// 	}
// }

// func BB() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before BB")
// 		c.Next()
// 		fmt.Println("After BB")
// 	}
// }

// func CC(c *gin.Context) {
// 	fmt.Println("Before CC")
// 	c.Next()
// 	fmt.Println("After CC")
// }

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// r.Use()
	// r.Use()
	// r.Use()

	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("v1/api")
	{
		MainGroup.GET("/checkStatus")
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.InitUserRouter(MainGroup)
		manageRouter.InitAdminRouter(MainGroup)
	}

	return r
}
