package main

import (
	"PieFireDire/config"
	"PieFireDire/handler"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.GetConf()

	e := echo.New()
	e.GET("/beef/summary", handler.GetBeefSummary)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%v:%v", conf.Host, conf.Port)))
}
