package services

import (
	"github.com/hongminhcbg/control-money/config"
	"github.com/hongminhcbg/control-money/daos"
	"github.com/hongminhcbg/control-money/dtos"
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
}

func NewUserService(conf *config.Config, userDao daos.UserDao) UserService {
	return &userServiceImpl{config: conf,
		userDao: userDao,
	}
}

func (service *userServiceImpl) Login(request dtos.LoginRequest) (*dtos.LoginResponse, error) {
	user, err := service.userDao.Login(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	response := dtos.LoginResponse{
		UserID:   user.ID,
		Username: user.Username,
		Name:     user.Name,
		Tocken:   "",
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
