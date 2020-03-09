package services

import (
	"github.com/hongminhcbg/control-money/config"
	"github.com/hongminhcbg/control-money/daos"
	"github.com/jinzhu/gorm"
)

type Provider interface {
	GetUserService() UserService
	GetAnalysisService() AnalysisService
}

type providerImpl struct {
	config *config.Config
	db *gorm.DB
}

func (provider *providerImpl) GetUserService() UserService {
	userDao := daos.NewUserDao(provider.db)
	return NewUserService(provider.config, userDao)
}

func (provider *providerImpl) GetAnalysisService() AnalysisService {
	return NewAnalysisService()
}

func NewProvider(conf *config.Config, db *gorm.DB) Provider {
	return &providerImpl{config: conf,db:db}
}
