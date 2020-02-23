package main

import (
	"github.com/hongminhcbg/control-money/config"
	"github.com/hongminhcbg/control-money/rounters"
)

func main()  {
	conf := &config.Config{}
	rounter := rounters.NewRounter(conf)
	app, _ := rounter.InitGin()
	app.Run(":8080")
}
