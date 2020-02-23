package services

import (
	"github.com/hongminhcbg/control-money/config"
)

type Provider interface {
	GetUserService() UserService
}

type providerImpl struct {
	config *config.Config
}

func (provider *providerImpl) GetUserService() UserService {
	return NewUserService(provider.config)
}

func NewProvider(conf *config.Config) Provider {
	return &providerImpl{config: conf,}
}
