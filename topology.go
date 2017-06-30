package main

import (
	"topology-go/config"
	"topology-go/controller"

	"github.com/labstack/echo"
)

func main() {
	log := config.NewLogger()
	log.Info("topology-go begin to work now :)")

	e := echo.New()
	e.GET("/", controller.RootGet)
	e.POST("/host", controller.HostPost)
	e.Logger.Fatal(e.Start(":9528"))
}
