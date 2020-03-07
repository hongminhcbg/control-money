package controlers

import (
	"github.com/gin-gonic/gin"

	"github.com/hongminhcbg/control-money/services"
)

type Controller struct {
	userService services.UserService
}

func NewController(provider services.Provider) Controller {
	return Controller{userService: provider.GetUserService(),}
}

func (ctl *Controller) Login(context *gin.Context) {
	context.JSON(400, gin.H{
		"message": "Service not support",
	})
}

func (ctl *Controller) Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Pong",
	})
}
