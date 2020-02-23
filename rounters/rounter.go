package rounters

import (
	"github.com/hongminhcbg/control-money/config"
	"github.com/hongminhcbg/control-money/services"
	"github.com/hongminhcbg/control-money/controlers"

	"github.com/gin-gonic/gin"
)

type Rounter struct {
	config *config.Config
}

func NewRounter(conf *config.Config) Rounter {
	return Rounter{config:conf,}
}

func (rounter *Rounter) InitGin() (*gin.Engine, error)  {
	providerService := services.NewProvider(rounter.config)
	controler := controlers.NewControler(providerService)

	engine := gin.Default()
	engine.GET("/ping", controler.Ping)
	return engine, nil
}