package services

import (
	"github.com/hongminhcbg/control-money/config"
	"github.com/hongminhcbg/control-money/daos"
	"github.com/hongminhcbg/control-money/middlewares"
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
	jwtClient := middlewares.NewJWT(provider.config.SecretKet)
	return NewUserService(provider.config, userDao, jwtClient)
}

func (provider *providerImpl) GetAnalysisService() AnalysisService {
	return NewAnalysisService()
}

func NewProvider(conf *config.Config, db *gorm.DB) Provider {
	return &providerImpl{config: conf,db:db}
}
