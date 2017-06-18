package core

import (
	"errors"
	"math"
	"sort"
	"topology-go/data"

	"github.com/emirpasic/gods/stacks/linkedliststack"
)

const shortGap = 1

// const longGap = 1

// Box store all data in a structured tree(itself)
type Box struct {
	width  int
	height int

	positionX int
	positionY int

	dataArray        []data.BaseData
	boxes            []Box
	heightWidthRatio float64
	records          *linkedliststack.Stack

	Key   string
	Value string
}

// NewBaseBox make box which only contains data.BaseData
func NewBaseBox(hwr float64, dataArray []data.BaseData) (Box, error) {
	sort.Sort(data.ByID(dataArray))
	box := Box{
		heightWidthRatio: hwr,
		dataArray:        dataArray,
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
		boxes:            boxes,
	}
	box.estimateSize()
	box.adaptBox()
	return box, nil
}

type byBoxSize []Box

func (s byBoxSize) Len() int      { return len(s) }
func (s byBoxSize) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byBoxSize) Less(i, j int) bool {
	if s[i].width != s[j].width {
		return s[i].width < s[j].width
	}
	return s[i].height < s[j].height
}

func (b *Box) shape() error {
	size := len(b.dataArray)
	if size == 0 {
		return errors.New("Empty Data")
	}
	floatSize := float64(size)
	b.width = int(math.Ceil(math.Sqrt(floatSize)))
	b.height = int(math.Ceil(floatSize / float64(b.width)))
	b.width = int(math.Ceil(floatSize / float64(b.height)))
	return nil
}

func (b *Box) setDataRelativePosition() error {
	if len(b.dataArray) > b.width*b.height {
		return errors.New("Data size error")
	}
	for index, baseData := range b.dataArray {
		baseData.RelativeX = index % b.width
		baseData.RelativeY = index / b.width
	}
	return nil
}

func (b *Box) estimateSize() {
	var estimateWeights float64
	for _, box := range b.boxes {
		estimateWeights += float64(box.width * box.height)
	}

	// add empty space
	estimateWeights *= 1.2
	estimateWeights = estimateWeights + 2*math.Sqrt(estimateWeights) + 1
	x := math.Sqrt(estimateWeights / b.heightWidthRatio)
	y := x * b.heightWidthRatio
	b.width = int(math.Floor(x))
	b.height = int(math.Floor(y))
}

func (b *Box) initRecords() {
	b.records = linkedliststack.New()
	record := data.NewRecord(0, 0, b.height)
	b.records.Push(record)
}

func (b *Box) adaptBox() {
	b.initRecords()
	check := true
	for _, box := range b.boxes {
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
	if b.enoughSpace(record.(data.Record), box) {
		b.fillBox(record.(data.Record), box)
	} else {
		b.addBox(box)
	}
	return true
}

func (b *Box) enoughSpace(record data.Record, box Box) bool {
	if record.PositionX+box.height > b.height {
		return false
	}
	if record.PositionY+box.width > b.width {
		return false
	}
	if record.PositionX+box.height > record.LimitHeight {
		return false
	}
	return true
}

func (b *Box) fillBox(record data.Record, box Box) {
	newRecord1 := data.NewRecord(record.PositionX+box.height+shortGap, record.PositionY, record.LimitHeight)
	newRecord2 := data.NewRecord(record.PositionX, record.PositionY+box.width+shortGap, record.PositionX+box.height+shortGap)
	if b.newRecordCheck(newRecord1) {
		b.records.Push(newRecord1)
	}
	if b.newRecordCheck(newRecord2) {
		b.records.Push(newRecord2)
	}
}

func (b *Box) newRecordCheck(record data.Record) bool {
	if record.PositionX >= b.height {
		return false
	}
	if record.PositionY >= b.width {
		return false
	}
	return true
}

func (b *Box) extendBox() {
	w := b.width + 1
	h := b.height + 1
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
	b.width = w
	b.height = h
}
