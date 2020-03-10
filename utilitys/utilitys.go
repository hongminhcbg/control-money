package utilitys

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/hongminhcbg/control-money/common"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/hongminhcbg/control-money/dtos"
)

type UserIdentity struct {
	UserID int64 `json:"id"`
}

// RequestHeader request header struct, get Bearer tocken
type RequestHeader struct {
	Rate   int    `header:"Rate"`
	Domain string `header:"Domain"`
	Tocken string `header:"Authorization"`
}

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

func parseBearerTocken(context *gin.Context) (*UserIdentity, error) {
	header := RequestHeader{}

	if err := context.BindHeader(&header); err != nil {
		return nil, err
	}
	// split header
	payloadArr := strings.Split(header.Tocken, " ")
	if len(payloadArr) != 2 {
		return nil, errors.New("parse error")
	}

	payload := payloadArr[1]
	// split payload header
	payloadArr = strings.Split(header.Tocken, ".")
	if len(payloadArr) != 3 {
		return nil, errors.New("parse error")
	}

	payload = payloadArr[1]
	// add "=" string to end base64 endcode
	if coutEqualStr := len(payload) % 4; coutEqualStr != 0 {
		for i := 0; i < (4 - coutEqualStr); i++ {
			payload += "="
		}
	}
	// decode base 64
	payloadByte, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return nil, err
	}

	// json Unmarshal
	userIdentity := UserIdentity{}
	if err := json.Unmarshal(payloadByte, &userIdentity); err != nil {
		return nil, err
	}

	if userIdentity.UserID <= 0 {
		return nil, errors.New("parse error")
	}

	return &userIdentity, nil
}

func SetUserID(context *gin.Context) error {
	UserIdentity, err := parseBearerTocken(context)
	if err != nil {
		return err
	}

	context.Set(common.HeaderUserID, UserIdentity.UserID)
	return nil
}

func GetUserID(context *gin.Context) (int64, error) {
	if val, existed := context.Get(common.HeaderUserID); existed {
		if valInt64, ok := val.(int64); ok {
			return valInt64, nil
		}
	}

	return 0, errors.New("get error")
}