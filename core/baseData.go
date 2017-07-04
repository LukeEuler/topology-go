package core

// BaseData represent the basic data for drawing
type BaseData struct {
	ID      int64             `json:"id"`
	Name    string            `json:"name"`
	TagMap  map[string]string `json:"tagMap"`
	TagList []string          `json:"tagList"`

	RelativeX int `json:"relative_x"`
	RelativeY int `json:"relative_y"`
	AbsoluteX int `json:"absolute_x"`
	AbsoluteY int `json:"absolute_y"`
}

func CompleteBaseData(data []*BaseData) []*BaseData {
	result := make([]*BaseData, 0, len(data))

	for _, baseData := range data {
		if len(baseData.Name) == 0 {
			continue
		}
		if baseData.TagMap == nil {
			baseData.TagMap = map[string]string{}
		}
		if baseData.TagList == nil {
			baseData.TagList = []string{}
		}

		result = append(result, baseData)
	}

	return result
}

// ByID support BaseData by ID
type ByID []*BaseData

func (s ByID) Len() int      { return len(s) }
func (s ByID) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByID) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}
