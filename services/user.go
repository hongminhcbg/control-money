package services

import (
	"github.com/hongminhcbg/control-money/config"
)

type UserService interface {
	Login() error
}

type userServiceImpl struct {
	config *config.Config
}

func NewUserService(conf *config.Config) UserService {
	return &userServiceImpl{config: conf,}
}

func (user *userServiceImpl) Login() error {
	return nil
}
