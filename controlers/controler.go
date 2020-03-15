package controlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hongminhcbg/control-money/common"
	"time"

	"github.com/hongminhcbg/control-money/dtos"
	"github.com/hongminhcbg/control-money/models"
	"github.com/hongminhcbg/control-money/services"
	"github.com/hongminhcbg/control-money/utilitys"
)

type Controller struct {
	userService services.UserService
	avrService  services.AverageService
}

func NewController(provider services.Provider) Controller {
	return Controller{userService: provider.GetUserService(),
		avrService: provider.GetAverageService()}
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
	if err != nil || logRequest.Money <= 0 {
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

func (ctl *Controller) GetLogs(context *gin.Context) {
	userID, err := utilitys.GetUserID(context)
	if err != nil {
		utilitys.ResponseError400(context, "get userID error")
		return
	}

	beginTime, err := stringYYYYMMDD2Time(context.Query(common.BeginTimeStampKey))
	if err != nil {
		utilitys.ResponseError400(context, "input format YYYY-MM-DD")
		return
	}

	endTime, err := stringYYYYMMDD2Time(context.Query(common.EndTimeStampKey))
	if err != nil {
		utilitys.ResponseError400(context, "input format YYYY-MM-DD")
		return
	}

	data, err := ctl.userService.GetLogsByTime(userID, beginTime, endTime)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {

		utilitys.ResponseSuccess200(context, data, "success")
	}
}

func stringYYYYMMDD2Time(input string) (*time.Time, error) {
	layout := "2006-01-02"
	result, err := time.Parse(layout,input)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ctl *Controller) AnalysisByTag(context *gin.Context) {
	userID, err := utilitys.GetUserID(context)
	if err != nil {
		utilitys.ResponseError400(context, "get userID error")
		return
	}

	beginTime, err := stringYYYYMMDD2Time(context.Query(common.BeginTimeStampKey))
	if err != nil {
		utilitys.ResponseError400(context, "input format YYYY-MM-DD")
		return
	}

	endTime, err := stringYYYYMMDD2Time(context.Query(common.EndTimeStampKey))
	if err != nil {
		utilitys.ResponseError400(context, "input format YYYY-MM-DD")
		return
	}

	data, err := ctl.userService.AnalysisByTag(userID, beginTime, endTime)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "success")
	}
}

func (ctl *Controller) AnalysisByDay(context *gin.Context) {
	userID, err := utilitys.GetUserID(context)
	if err != nil {
		utilitys.ResponseError400(context, "get userID error")
		return
	}

	beginTime, err := stringYYYYMMDD2Time(context.Query(common.BeginTimeStampKey))
	if err != nil {
		utilitys.ResponseError400(context, "input format YYYY-MM-DD")
		return
	}

	endTime, err := stringYYYYMMDD2Time(context.Query(common.EndTimeStampKey))
	if err != nil {
		utilitys.ResponseError400(context, "input format YYYY-MM-DD")
		return
	}

	data, err := ctl.userService.AnalysisByDay(userID, beginTime, endTime)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		utilitys.ResponseSuccess200(context, data, "input format YYYY-MM-DD")
	}
}

func (ctl *Controller)GetAverageByDay(context *gin.Context)  {
	userID, err := utilitys.GetUserID(context)
	if err != nil {
		utilitys.ResponseError400(context, "get userID error")
		return
	}

	beginTime, err := stringYYYYMMDD2Time(context.Query(common.BeginTimeStampKey))
	if err != nil {
		utilitys.ResponseError400(context, "input format YYYY-MM-DD")
		return
	}

	endTime, err := stringYYYYMMDD2Time(context.Query(common.EndTimeStampKey))
	if err != nil {
		utilitys.ResponseError400(context, "input format YYYY-MM-DD")
		return
	}

	money, err := ctl.avrService.CalculateByDay(userID, beginTime, endTime)
	if err != nil {
		utilitys.ResponseError400(context, err.Error())
	} else {
		result := dtos.AverageMoneyPerDayResponse{Money: money}
		utilitys.ResponseSuccess200(context, result, "success")
	}
}

func (ctl *Controller) Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Pong",
	})
}
