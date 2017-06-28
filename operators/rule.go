package operators

import (
	"errors"

	"topology-go/core"
)

type RuleConfig struct {
	HeightWidthRatio float64
	Tag              []string
	ShowEmptyValue   bool
	EmptyValueName   string
}

func defaultRuleConfig() RuleConfig {
	return RuleConfig{
		HeightWidthRatio: 1,
		Tag:              []string{},
		ShowEmptyValue:   false,
	}
}

func Rule(dataArray []core.BaseData) (core.Box, error) {
	return RuleWithConfig(dataArray, defaultRuleConfig())
}

func RuleWithConfig(dataArray []core.BaseData, conf RuleConfig) (result core.Box, err error) {
	if dataArray == nil || len(dataArray) == 0 {
		return result, errors.New("empty data")
	}

	if conf.HeightWidthRatio <= 0 {
		conf.HeightWidthRatio = 1
	}

	if conf.Tag == nil {
		conf.Tag = []string{}
	}

	return rule(dataArray, conf)
}

func rule(dataArray []core.BaseData, conf RuleConfig) (result core.Box, err error) {
	// TODO too complicate, no test
	if len(conf.Tag) == 0 {
		return core.NewBaseBox(conf.HeightWidthRatio, dataArray)
	} else {
		tag := conf.Tag[0]
		midResult := make(map[string][]core.BaseData)
		for _, dataPoint := range dataArray {
			tagValue, ok := dataPoint.TagMap[tag]
			if conf.ShowEmptyValue {
				if !ok {
					tagValue = conf.EmptyValueName
				}
				_, okMid := midResult[tagValue]
				if !okMid {
					midResult[tagValue] = []core.BaseData{dataPoint}
				} else {
					midResult[tagValue] = append(midResult[tagValue], dataPoint)
				}
			} else {
				if !ok {
					continue
				}
				_, okMid := midResult[tagValue]
				if !okMid {
					midResult[tagValue] = []core.BaseData{dataPoint}
				} else {
					midResult[tagValue] = append(midResult[tagValue], dataPoint)
				}
			}
		}

		conf.Tag = conf.Tag[1:]
		boxes := make([]core.Box, 0)
		for tagValue, subDataArray := range midResult {
			box, err := rule(subDataArray, conf)
			if err != nil {
				continue
			}
			box.Key = tag
			box.Value = tagValue
			boxes = append(boxes, box)
		}

		if len(boxes) == 0 {
			return result, errors.New("empty boxes")
		}
		return core.NewAdvanceBox(conf.HeightWidthRatio, boxes)
	}
}
