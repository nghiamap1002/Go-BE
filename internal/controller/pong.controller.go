package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongCotroller() *PongController {
	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {
	fmt.Println("Pong handler")
	name := c.DefaultQuery("name", "concobebe")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "ping...pong" + name,
		"uid":     uid,
		"users":   []string{"abc", "concho", "conmeo"},
	})
}
