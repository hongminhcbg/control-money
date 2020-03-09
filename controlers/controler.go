package controlers

import (
	"github.com/gin-gonic/gin"

	"github.com/hongminhcbg/control-money/dtos"
	"github.com/hongminhcbg/control-money/models"
	"github.com/hongminhcbg/control-money/services"
	"github.com/hongminhcbg/control-money/utilitys"
)

type Controller struct {
	userService services.UserService
}

func NewController(provider services.Provider) Controller {
	return Controller{userService: provider.GetUserService(),}
}

func (ctl *Controller) Login(context *gin.Context) {
	var request dtos.LoginRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		utilitys.ResponseError400(context, "login error")
		return
	}

	data, err := ctl.userService.Login(request)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "login success")
	}
}

func (ctl *Controller) CreateUser(context *gin.Context) {
	var request dtos.CreateUserRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
		return
	}

	acc := models.User{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
	}
	data, err := ctl.userService.Create(acc)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "success")
	}
}

func (ctl *Controller) CreateLog(context *gin.Context) {

}

func (ctl *Controller) Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Pong",
	})
}
