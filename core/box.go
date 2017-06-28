package core

import (
	"errors"
	"math"
	"sort"

	"github.com/emirpasic/gods/stacks/linkedliststack"
)

const narrowGap = 1

// const longGap = 1

// Box store all data in a structured tree(itself)
type Box struct {
	Width  int `json:"width"`
	Height int `json:"height"`

	PositionX int `json:"position_x"`
	PositionY int `json:"position_y"`

	DataArray        []BaseData `json:"dataArray,omitempty"`
	Boxes            []Box `json:"boxes,omitempty"`
	heightWidthRatio float64
	records          *linkedliststack.Stack

	Key   string `json:"tagKey,omitempty"`
	Value string `json:"tagValue,omitempty"`
}

// NewBaseBox make box which only contains data.BaseData
func NewBaseBox(hwr float64, dataArray []BaseData) (Box, error) {
	sort.Sort(ByID(dataArray))
	box := Box{
		heightWidthRatio: hwr,
		DataArray:        dataArray,
	}
	err := box.shape()
	if err != nil {
		return box, err
	}
	err = box.setDataRelativePosition()
	if err != nil {
		return box, err
	}
	return box, nil
}

// NewAdvanceBox make box which only contains data.BaseData
func NewAdvanceBox(hwr float64, boxes []Box) (Box, error) {
	sort.Sort(byBoxSize(boxes))
	box := Box{
		heightWidthRatio: hwr,
		Boxes:            boxes,
	}
	box.estimateSize()
	box.adaptBox()
	return box, nil
}

type byBoxSize []Box

func (s byBoxSize) Len() int      { return len(s) }
func (s byBoxSize) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byBoxSize) Less(i, j int) bool {
	if s[i].Width != s[j].Width {
		return s[i].Width < s[j].Width
	}
	return s[i].Height < s[j].Height
}

func (b *Box) shape() error {
	size := len(b.DataArray)
	if size == 0 {
		return errors.New("Empty Data")
	}
	floatSize := float64(size)
	b.Width = int(math.Ceil(math.Sqrt(floatSize)))
	b.Height = int(math.Ceil(floatSize / float64(b.Width)))
	b.Width = int(math.Ceil(floatSize / float64(b.Height)))
	return nil
}

func (b *Box) setDataRelativePosition() error {
	if len(b.DataArray) > b.Width*b.Height {
		return errors.New("Data size error")
	}
	for index, baseData := range b.DataArray {
		baseData.RelativeX = index % b.Width
		baseData.RelativeY = index / b.Width
	}
	return nil
}

func (b *Box) estimateSize() {
	var estimateWeights float64
	for _, box := range b.Boxes {
		estimateWeights += float64(box.Width * box.Height)
	}

	// add empty space
	estimateWeights *= 1.2
	estimateWeights = estimateWeights + 2*math.Sqrt(estimateWeights) + 1
	x := math.Sqrt(estimateWeights / b.heightWidthRatio)
	y := x * b.heightWidthRatio
	b.Width = int(math.Floor(x))
	b.Height = int(math.Floor(y))
}

func (b *Box) initRecords() {
	b.records = linkedliststack.New()
	record := NewRecord(0, 0, b.Height)
	b.records.Push(record)
}

func (b *Box) adaptBox() {
	b.initRecords()
	check := true
	for _, box := range b.Boxes {
		check = b.addBox(box)
		if !check {
			break
		}
	}
	if check {
		return
	}
	b.extendBox()
	b.adaptBox()
}

func (b *Box) addBox(box Box) bool {
	record, ok := b.records.Pop()
	if !ok {
		return false
	}
	if b.enoughSpace(record.(Record), box) {
		b.fillBox(record.(Record), box)
	} else {
		b.addBox(box)
	}
	return true
}

func (b *Box) enoughSpace(record Record, box Box) bool {
	if record.PositionX+box.Height > b.Height {
		return false
	}
	if record.PositionY+box.Width > b.Width {
		return false
	}
	if record.PositionX+box.Height > record.LimitHeight {
		return false
	}
	return true
}

func (b *Box) fillBox(record Record, box Box) {
	newRecord1 := NewRecord(record.PositionX+box.Height+narrowGap, record.PositionY, record.LimitHeight)
	newRecord2 := NewRecord(record.PositionX, record.PositionY+box.Width+narrowGap, record.PositionX+box.Height+narrowGap)
	if b.newRecordCheck(newRecord1) {
		b.records.Push(newRecord1)
	}
	if b.newRecordCheck(newRecord2) {
		b.records.Push(newRecord2)
	}
}

func (b *Box) newRecordCheck(record Record) bool {
	if record.PositionX >= b.Height {
		return false
	}
	if record.PositionY >= b.Width {
		return false
	}
	return true
}

func (b *Box) extendBox() {
	w := b.Width + 1
	h := b.Height + 1
	p0 := float64(h)/float64(w) - b.heightWidthRatio
	if p0 > 0 {
		p1 := float64(h)/float64(w+1) - b.heightWidthRatio
		for math.Abs(p1) <= math.Abs(p0) {
			p0 = p1
			w++
			p1 = float64(h)/float64(w+1) - b.heightWidthRatio
		}
	} else if p0 < 0 {
		p1 := float64(h+1)/float64(w) - b.heightWidthRatio
		for math.Abs(p1) <= math.Abs(p0) {
			p0 = p1
			h++
			p1 = float64(h+1)/float64(w) - b.heightWidthRatio
		}
	}
	b.Width = w
	b.Height = h
}
