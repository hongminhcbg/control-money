package utilitys

import (
	"github.com/gin-gonic/gin"

	"github.com/hongminhcbg/control-money/dtos"
)

// MakeRespond make respond when frontend call API
func makeResponse(data interface{}, code int, msg string) dtos.Response {
	return dtos.Response{
		Data: data,
		Meta: struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{code, msg},
	}
}

func Response(context *gin.Context, data interface{}, code int, msg string)  {
	resp := makeResponse(data, code, msg)
	context.JSON(code, resp)
}

func ResponseError400(context *gin.Context, msg string)  {
	Response(context, nil, 400, msg)
}

func ResponseSuccess200(context *gin.Context, data interface{}, msg string)  {
	Response(context, data, 200, msg)
}
