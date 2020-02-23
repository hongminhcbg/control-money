package controlers

import (
	"github.com/gin-gonic/gin"

	"github.com/hongminhcbg/control-money/services"
)

type Controler struct {
	userService services.UserService
}

func NewControler(provider services.Provider) Controler {
	return Controler{userService: provider.GetUserService(),}
}

func (ctl *Controler) Login(context *gin.Context) {
	context.JSON(400, gin.H{
		"message": "Service not support",
	})
}

func (ctl *Controler) Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Pong",
	})
}
