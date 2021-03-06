package core

// Record represent a point where a new b may be placed
type Record struct {
	PositionX   int
	PositionY   int
	LimitHeight int
}

// NewRecord create new Record
func NewRecord(x, y, height int) Record {
	return Record{
		PositionX:   x,
		PositionY:   y,
		LimitHeight: height,
	}
}
