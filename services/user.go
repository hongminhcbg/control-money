package services

import (
	"github.com/hongminhcbg/control-money/config"
	"github.com/hongminhcbg/control-money/dtos"
	"github.com/hongminhcbg/control-money/models"
)

type UserService interface {
	Login(user models.User) error
	Create(user models.User) (*models.User, error)
	CreateLog(log models.Log) (*models.Log, error)
	GetAverageMonth(request dtos.AvrMoneyRequest) (*dtos.AvrMoneyResponse, error)
}

type userServiceImpl struct {
	config *config.Config
}

func NewUserService(conf *config.Config) UserService {
	return &userServiceImpl{config: conf,}
}

func (service *userServiceImpl) Login(user models.User) error {
	return nil
}

func (service *userServiceImpl) Create(user models.User) (*models.User, error) {
	return nil, nil
}

func (service *userServiceImpl) CreateLog(log models.Log) (*models.Log, error) {
	return nil, nil
}

func (service *userServiceImpl) GetAverageMonth(request dtos.AvrMoneyRequest) (*dtos.AvrMoneyResponse, error) {
	return nil, nil
}
