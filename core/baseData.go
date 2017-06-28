package core

// BaseData represent the basic data for drawing
type BaseData struct {
	ID      int64
	name    string
	TagMap  map[string]string
	TagList []string

	RelativeX int
	RelativeY int

	absoluteX int
	absoluteY int
}

// ByID support BaseData by ID
type ByID []BaseData

func (s ByID) Len() int      { return len(s) }
func (s ByID) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByID) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}
