package rounters

import (
	"github.com/gin-gonic/gin"
	"github.com/hongminhcbg/control-money/utilitys"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/contrib/jwt"

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

	accountAuthMiddleWare := middlewares.CheckAPIKey{ApiKey: router.config.APIKey}
	{
		account := engine.Group("/api/v1/account")
		account.Use(accountAuthMiddleWare.Check)
		account.POST("", controller.CreateUser)
		account.POST("login", controller.Login)
	}

	{
		log := engine.Group("/api/v1/log")
		log.Use(jwt.Auth(router.config.SecretKet))
		log.Use(middlewares.SetUserID)
		log.POST("/ping", controller.Ping)
	}

	return engine, nil
}