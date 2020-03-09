package services

import (
	"github.com/hongminhcbg/control-money/config"
	"github.com/hongminhcbg/control-money/daos"
	"github.com/hongminhcbg/control-money/dtos"
	"github.com/hongminhcbg/control-money/middlewares"
	"github.com/hongminhcbg/control-money/models"
)

type UserService interface {
	Login(request dtos.LoginRequest) (*dtos.LoginResponse, error)
	Create(user models.User) (*models.User, error)
	CreateLog(log models.Log) (*models.Log, error)
	GetAverageMonth(request dtos.AvrMoneyRequest) (*dtos.AvrMoneyResponse, error)
}

type userServiceImpl struct {
	config  *config.Config
	userDao daos.UserDao
	jwt     middlewares.JWT
}

func NewUserService(conf *config.Config, userDao daos.UserDao, jwt middlewares.JWT) UserService {
	return &userServiceImpl{config: conf,
		userDao: userDao,
		jwt:jwt,
	}
}

func (service *userServiceImpl) Login(request dtos.LoginRequest) (*dtos.LoginResponse, error) {
	user, err := service.userDao.Login(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	tocken, err := service.jwt.CreateTocken(user.ID)
	if err != nil {
		return nil, err
	}

	response := dtos.LoginResponse{
		UserID:   user.ID,
		Username: user.Username,
		Name:     user.Name,
		Tocken:   tocken,
		Money:    user.Money,
	}
	return &response, nil
}

func (service *userServiceImpl) Create(user models.User) (*models.User, error) {
	return service.userDao.Create(user)
}

func (service *userServiceImpl) CreateLog(log models.Log) (*models.Log, error) {
	return nil, nil
}

func (service *userServiceImpl) GetAverageMonth(request dtos.AvrMoneyRequest) (*dtos.AvrMoneyResponse, error) {
	return nil, nil
}
