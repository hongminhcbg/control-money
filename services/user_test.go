package services

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/hongminhcbg/control-money/config"
	"github.com/hongminhcbg/control-money/daos"
	"github.com/hongminhcbg/control-money/middlewares"
	"github.com/hongminhcbg/control-money/models"

)

const dbLocalURL = "root:1@tcp(localhost:3306)/cm?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"

func Test_userServiceImpl_Create(t *testing.T) {
	db, err := gorm.Open("mysql", dbLocalURL)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = db.Close()
	}()

	type fields struct {
		config  *config.Config
		userDao daos.UserDao
		jwt     middlewares.JWT
	}
	type args struct {
		user models.User
	}

	userNameRandom := uuid.New().String()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "normal case",
			fields: fields{
				config:  nil,
				userDao: daos.NewUserDao(db),
				jwt:     nil,
			},
			args: args{
				user: models.User{
					Name:     "hong",
					Username: userNameRandom,
					Money:    100,
				},
			},
			want: &models.User{
				Name:     "hong",
				Username: userNameRandom,
				Money:    100,
			},
			wantErr: false,
		},
		{
			name: "duplicate user name",
			fields: fields{
				config:  nil,
				userDao: daos.NewUserDao(db),
				jwt:     nil,
			},
			args: args{
				user: models.User{
					Name:     "hong",
					Username: userNameRandom,
					Money:    100,
				},
			},
			want: nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &userServiceImpl{
				config:  tt.fields.config,
				userDao: tt.fields.userDao,
				jwt:     tt.fields.jwt,
			}
			got, err := service.Create(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && tt.want != nil {
				tt.want.ID = got.ID
				tt.want.UpdateAt = got.UpdateAt
				tt.want.CreateAt = got.CreateAt
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}