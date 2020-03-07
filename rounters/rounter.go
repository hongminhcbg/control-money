package rounters

import (
	"github.com/gin-gonic/gin"

	"github.com/hongminhcbg/control-money/config"
	"github.com/hongminhcbg/control-money/controlers"
	"github.com/hongminhcbg/control-money/services"

)

type Router struct {
	config *config.Config
}

func NewRouter(conf *config.Config) Router {
	return Router{config: conf,}
}

func (router *Router) InitGin() (*gin.Engine, error)  {
	providerService := services.NewProvider(router.config)
	controller := controlers.NewController(providerService)

	engine := gin.Default()
	engine.GET("/ping", controller.Ping)
	return engine, nil
}