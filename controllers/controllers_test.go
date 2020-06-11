package controllers

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/hongminhcbg/control-money/models"
	"github.com/hongminhcbg/control-money/services"
)

func TestController_CreateUser(t *testing.T) {
	type fields struct {
		userService services.UserService
		avrService  services.AverageService
	}
	type args struct {
		context *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "test case 1",
			fields: fields{userService: nil, avrService: nil},
			args:   args{nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := &Controller{
				userService: tt.fields.userService,
				avrService:  tt.fields.avrService,
			}

			user := models.User{
				Name:     "minh",
				Username: "nguyen hong",
				Money:    0,
				Password: "",
			}

			_, err := ctl.userService.Create(user)

			assert.NotNil(t, err, "create user false because service is nil")
		})
	}
}