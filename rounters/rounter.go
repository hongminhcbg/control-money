package rounters

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/hongminhcbg/control-money/config"
	"github.com/hongminhcbg/control-money/controlers"
	"github.com/hongminhcbg/control-money/middlewares"
	"github.com/hongminhcbg/control-money/services"

)

type Router struct {
	config *config.Config
	db *gorm.DB
}

func NewRouter(conf *config.Config, db *gorm.DB) Router {
	return Router{config: conf, db:db}
}

func (router *Router) InitGin() (*gin.Engine, error)  {


	providerService := services.NewProvider(router.config, router.db)
	controller := controlers.NewController(providerService)

	engine := gin.Default()
	engine.Use(middlewares.CORSMiddleware())
	engine.GET("/ping", controller.Ping)

	accountAuth := middlewares.CheckAPIKey{ApiKey:router.config.APIKey}
	{
		account := engine.Group("/api/v1/account")
		account.Use(accountAuth.Check)
		account.POST("", controller.CreateUser)
	}


	return engine, nil
}