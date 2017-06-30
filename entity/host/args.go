package host

type Args struct {
	HeightWidthRatio float64 `json:"heightWidthRatio"`
	Filter           filter `json:"filter"`
	Tag              []string `json:"tag"`
	ShowEmptyValue   bool `json:"showEmptyValue"`
	EmptyValueName   string `json:"emptyValueName"`
}

type filter struct {
	List []string `json:"list"`
	Map  map[string]string `json:"map"`
}

func NewArgs() *Args {
	return new(Args)
}
