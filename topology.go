package main

import (
	"encoding/json"

	"topology-go/core"
	"topology-go/operators"
	"topology-go/resources"
)

func main() {
	data := resources.GetBaseData("./resources/data.json")
	data = core.CompleteBaseData(data)
	filterTagList := []string{}
	filterTagMap := map[string]string{}
	data = operators.Filter(filterTagList, filterTagMap, data)
	conf := operators.RuleConfig{
		HeightWidthRatio: 0.75,
		Tag:              []string{"role"},
		ShowEmptyValue:   true,
		EmptyValueName:   "Not Found",
	}
	box, err := operators.RuleWithConfig(data, conf)
	if err != nil {
		println(err.Error())
		return
	}
	box.SetPosition()
	bb, err := json.Marshal(box)
	if err != nil {
		println(err.Error())
		return
	}
	println(string(bb))
}
