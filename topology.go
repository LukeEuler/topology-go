package main

import (
	"github.com/labstack/echo"
	"github.com/LukeEuler/topology-go/config"
	"github.com/LukeEuler/topology-go/controller"
)

func main() {
	log := config.NewLogger()
	log.Info("topology-go begin to work now :)")

	e := echo.New()
	e.GET("/", controller.RootGet)
	e.POST("/host", controller.HostPost)
	e.Logger.Fatal(e.Start(":9528"))
}
