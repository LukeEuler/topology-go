package controller

import (
	"net/http"

	"github.com/LukeEuler/topology-go/config"
	"github.com/LukeEuler/topology-go/core"
	"github.com/LukeEuler/topology-go/entity/host"
	"github.com/LukeEuler/topology-go/operators"
	"github.com/LukeEuler/topology-go/resources"
	"github.com/labstack/echo"
)

// HostPost deal with query
func HostPost(c echo.Context) (err error) {
	args := host.NewArgs()
	if err = c.Bind(args); err != nil {
		return
	}
	box := getHost(args)
	return c.JSON(http.StatusOK, box)
}

func getHost(args *host.Args) *core.Box {
	data := resources.GetBaseData("./resources/data.json")
	data = core.CompleteBaseData(data)
	data = operators.Filter(args.Filter.List, args.Filter.Map, data)

	conf := operators.RuleConfig{
		HeightWidthRatio: args.HeightWidthRatio,
		Tag:              args.Tag,
		ShowEmptyValue:   args.ShowEmptyValue,
		EmptyValueName:   args.EmptyValueName,
	}

	box, err := operators.RuleWithConfig(data, conf)
	if err != nil {
		log := config.NewLogger()
		log.Error(err.Error())
		return nil
	}
	box.SetPosition()
	return box
}
