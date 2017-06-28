package resources

import (
	"io/ioutil"

	"encoding/json"
	"fmt"
	"os"

	"topology-go/core"
)

func GetBaseData(path string) []core.BaseData {
	raw, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []core.BaseData
	json.Unmarshal(raw, &c)
	return c
}
