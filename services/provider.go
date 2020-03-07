package services

import (
	"github.com/hongminhcbg/control-money/config"
)

type Provider interface {
	GetUserService() UserService
	GetAnalysisService() AnalysisService
}

type providerImpl struct {
	config *config.Config
}

func (provider *providerImpl) GetUserService() UserService {
	return NewUserService(provider.config)
}

func (provider *providerImpl) GetAnalysisService() AnalysisService {
	return NewAnalysisService()
}

func NewProvider(conf *config.Config) Provider {
	return &providerImpl{config: conf,}
}
