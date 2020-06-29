package daos

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/hongminhcbg/control-money/bcrypt"
	"github.com/hongminhcbg/control-money/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const mySQLLocal = "root:bW90aGVyIGZ1Y2tlciBub29i@tcp(localhost:3307)/cm?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"

func Test_userDaoImpl_Create(t *testing.T) {
	db, err := gorm.Open("mysql", mySQLLocal)
	if err != nil {
		panic(err)
	}

	bcryptClient := bcrypt.NewBcryptClient()
	userNameRandom := uuid.New().String()

	type fields struct {
		db           *gorm.DB
		bcryptClient bcrypt.BcryptClient
	}
	type args struct {
		user models.User
	}
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
				db:           db,
				bcryptClient: bcryptClient,
			},
			args: args{
				models.User{
					Name:     "hongminh234",
					Password: "229297",
					Username: userNameRandom,
				},
			},
			wantErr: false,
			want:    nil,
		},
		{
			name: "douplicate userName",
			fields: fields{
				db:           db,
				bcryptClient: bcryptClient,
			},
			args: args{
				models.User{
					Name:     "nguyen hong ming",
					Password: "fdsjfsdfsd",
					Username: userNameRandom,
				},
			},
			wantErr: true,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := &userDaoImpl{
				db:           tt.fields.db,
				bcryptClient: tt.fields.bcryptClient,
			}
			got, err := dao.Create(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userDaoImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("userDaoImpl.Create() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_userDaoImpl_Login(t *testing.T) {
	db, err := gorm.Open("mysql", mySQLLocal)
	if err != nil {
		panic(err)
	}

	bcryptClient := bcrypt.NewBcryptClient()
	type fields struct {
		db           *gorm.DB
		bcryptClient bcrypt.BcryptClient
	}
	type args struct {
		userName string
		pass     string
	}
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
				db:           db,
				bcryptClient: bcryptClient,
			},
			args: args{
				userName: "hongminhnguyen",
				pass:     "229297",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "wrong password",
			fields: fields{
				db:           db,
				bcryptClient: bcryptClient,
			},
			args: args{
				userName: "hongminhnguyen",
				pass:     "229297777",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "user name not existed",
			fields: fields{
				db:           db,
				bcryptClient: bcryptClient,
			},
			args: args{
				userName: "xxxxxxxxxxx",
				pass:     "229297777",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := &userDaoImpl{
				db:           tt.fields.db,
				bcryptClient: tt.fields.bcryptClient,
			}
			got, err := dao.Login(tt.args.userName, tt.args.pass)
			if (err != nil) != tt.wantErr {
				t.Errorf("userDaoImpl.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("userDaoImpl.Login() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
