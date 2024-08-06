package routers

import (
	"personal/ShopDev/Go-BE/internal/routers/manage"
	"personal/ShopDev/Go-BE/internal/routers/user"
)

type RouterGroup struct {
	User   user.UserRouterGroup
	Manage manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
