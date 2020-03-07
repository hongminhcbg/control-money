package main

import (
	"github.com/hongminhcbg/control-money/config"
	"github.com/hongminhcbg/control-money/rounters"
)

func main()  {
	conf := &config.Config{}
	router := rounters.NewRouter(conf)
	app, _ := router.InitGin()
	app.Run(":8080")
}
