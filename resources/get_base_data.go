package resources

import (
	"io/ioutil"

	"encoding/json"
	"fmt"
	"os"

	"topology-go/core"
)

func GetBaseData() []core.BaseData {
	raw, err := ioutil.ReadFile("./data.json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []core.BaseData
	json.Unmarshal(raw, &c)
	return c
}
