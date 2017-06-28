package operators

import (
	"topology-go/core"

	set "github.com/deckarep/golang-set"
)

// Filter remove the data without special tags
func Filter(tagList []string, tagMap map[string]string, dataArray []core.BaseData) {
	if len(tagList) != 0 {
		dataArray = filterTagList(tagList, dataArray)
	}
	if tagMap != nil && len(tagMap) != 0 {
		dataArray = filterTagMap(tagMap, dataArray)
	}
}

func filterTagList(tagList []string, dataArray []core.BaseData) (result []core.BaseData) {
	filterTagSet := makeSet(tagList)
	for _, baseData := range dataArray {
		dataTagSet := makeSet(baseData.TagList)
		if filterTagSet.IsSubset(dataTagSet) {
			result = append(result, baseData)
		}
	}

	return result
}

func makeSet(strList []string) set.Set {
	result := set.NewSet()
	for _, str := range strList {
		result.Add(str)
	}
	return result
}

func filterTagMap(tagMap map[string]string, dataArray []core.BaseData) (result []core.BaseData) {
	for _, baseData := range dataArray {
		if mapContains(tagMap, baseData.TagMap) {
			result = append(result, baseData)
		}
	}
	return result
}

func mapContains(unionMap map[string]string, subMap map[string]string) bool {
	if len(subMap) == 0 {
		if len(unionMap) == 0 {
			return true
		}
		return false

	}
	for key, value := range subMap {
		unionMapValue, ok := unionMap[key]
		if !ok {
			return false
		}
		if value != unionMapValue {
			return false
		}
	}
	return true
}
