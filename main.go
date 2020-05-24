package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kelseyhightower/envconfig"

	"github.com/hongminhcbg/control-money/config"
	"github.com/hongminhcbg/control-money/rounters"
)

func main() {
	conf := config.Config{}
	if err := envconfig.Process("", &conf); err != nil {
		panic(err)
	}
	db, err := gorm.Open("mysql", conf.MySQLURL)
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic("open db error: " + err.Error())
	}

	if err := db.DB().Ping(); err != nil {
		panic("ping db error: " + err.Error())
	}
	router := rounters.NewRouter(&conf, db)
	app, _ := router.InitGin()
	_ = app.Run(":80")
}
