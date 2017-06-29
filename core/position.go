package core

func (b *Box) SetPosition() {
	setDataPosition(b.AbsoluteX, b.AbsoluteY, b.DataPoints)
	setBoxPosition(b.AbsoluteX, b.AbsoluteY, b.Boxes)
}

func setBoxPosition(x, y int, boxes []*Box) {
	if boxes == nil {
		return
	}

	for _, box := range boxes {
		box.AbsoluteX = x + box.RelativeX
		box.AbsoluteY = y + box.RelativeY
		setDataPosition(box.AbsoluteX, box.AbsoluteY, box.DataPoints)
		setBoxPosition(box.AbsoluteX, box.AbsoluteY, box.Boxes)
	}
}

func setDataPosition(x, y int, points []*BaseData) {
	if points == nil {
		return
	}

	for _, point := range points {
		point.AbsoluteX = x + point.RelativeX
		point.AbsoluteY = y + point.RelativeY
	}
}
