package controlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hongminhcbg/control-money/common"
	"strconv"
	"time"

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
	userID, err := utilitys.GetUserID(context)
	if err != nil {
		utilitys.ResponseError400(context, "get userid error")
		return
	}

	var logRequest dtos.CreateLogRequest
	err = context.ShouldBindJSON(&logRequest)
	if err != nil {
		utilitys.ResponseError400(context, "json decode error")
		return
	}

	log := models.Log{
		Detail: logRequest.Detail,
		Money:  logRequest.Money,
		Tag:    logRequest.Tag,
		UserID: userID,
	}

	data, err := ctl.userService.CreateLog(log)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "create log success")
	}

}

func string2Time(input string) (*time.Time, error)  {
	timeSecondInt, err := strconv.Atoi(input)
	if err != nil {
		return nil, err
	}

	t := time.Unix(int64(timeSecondInt), 0)
	return &t, err
}

func (ctl *Controller)AnalysisByTag(context *gin.Context)  {
	userID, err :=  utilitys.GetUserID(context)
	if err != nil {
		utilitys.ResponseError400(context, "get userID error")
		return
	}

	beginTime, err := string2Time(context.Query(common.BeginTimeStampKey))
	if err != nil {
		utilitys.ResponseError400(context, "parse time error")
		return
	}

	endTime, err := string2Time(context.Query(common.EndTimeStampKey))
	if err != nil {
		utilitys.ResponseError400(context, "parse time error")
		return
	}

	data, err := ctl.userService.AnalysisByTag(userID, beginTime, endTime)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "success")
	}
}

func (ctl *Controller) Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Pong",
	})
}
