package flynex

import (
	fyne "fyne.io/fyne/v2"
)

type renderer struct {
	container *Container
}

func (r *renderer) unmarshalPos(pos fyne.Position) (main, cross float32) {
	if r.container.column {
		return pos.Y, pos.X
	}
	return pos.X, pos.Y
}

func (r *renderer) marshalPos(main, cross float32) fyne.Position {
	if r.container.column {
		return fyne.NewPos(cross, main)
	}
	return fyne.NewPos(main, cross)
}

func (r *renderer) unmarshalSize(size fyne.Size) (main, cross float32) {
	if r.container.column {
		return size.Height, size.Width
	}
	return size.Width, size.Height
}

func (r *renderer) marshalSize(main, cross float32) fyne.Size {
	if r.container.column {
		return fyne.NewSize(cross, main)
	}
	return fyne.NewSize(main, cross)
}

func (r *renderer) Layout(size fyne.Size) {
	var mainAxisSizedElementsTotalSizes float32
	var totalFlexes float32

	for _, item := range r.container.Items {
		main, cross := r.unmarshalSize(item.MinSize())
		if item.flex != 0 {
			totalFlexes += float32(item.flex)
		} else {
			if item.mainSize != 0 {
				main = item.mainSize
			}
			item.Resize(r.marshalSize(main, cross))
			mainAxisSizedElementsTotalSizes += main
		}
	}

	maxMainSize, maxCrossSize := r.unmarshalSize(size)
	remainMainAxisSizes := maxMainSize - mainAxisSizedElementsTotalSizes
	if remainMainAxisSizes <= 0 {
		var mainAcc float32
		for _, item := range r.container.Items {
			mainSize, _ := r.unmarshalSize(item.MinSize())
			_, crossPos := r.unmarshalPos(item.Position())
			item.Move(r.marshalPos(mainAcc, crossPos))
			mainAcc += mainSize
		}
		return
	} else if totalFlexes > 0 {
		sizePerFlex := remainMainAxisSizes / totalFlexes
		for _, item := range r.container.Items {
			if item.flex != 0 {
				_, cross := r.unmarshalSize(item.MinSize())
				item.Resize(r.marshalSize(sizePerFlex*float32(item.flex), cross))
			}
		}
		remainMainAxisSizes = 0
	}

	switch r.container.justify {
	case justifySpaceBetween:
		if len(r.container.Items) == 1 {
			return
		}
		slots := len(r.container.Items) - 1
		slotSize := remainMainAxisSizes / float32(slots)
		var acc float32
		for _, item := range r.container.Items {
			mainSize, _ := r.unmarshalSize(item.Size())
			_, crossPos := r.unmarshalPos(item.Position())
			item.Move(r.marshalPos(acc, crossPos))
			acc += mainSize + slotSize
		}
	case justifySpaceEvenly:
		slots := len(r.container.Items) + 1
		slotSize := remainMainAxisSizes / float32(slots)
		acc := slotSize
		for _, item := range r.container.Items {
			mainSize, _ := r.unmarshalSize(item.Size())
			_, crossPos := r.unmarshalPos(item.Position())
			item.Move(r.marshalPos(acc, crossPos))
			acc += mainSize + slotSize
		}
	case justifyFlexStart:
		var acc float32
		for _, item := range r.container.Items {
			mainSize, _ := r.unmarshalSize(item.Size())
			_, crossPos := r.unmarshalPos(item.Position())
			item.Move(r.marshalPos(acc, crossPos))
			acc += mainSize
		}
	case justifyFlexEnd:
		acc := remainMainAxisSizes
		for _, item := range r.container.Items {
			mainSize, _ := r.unmarshalSize(item.Size())
			_, crossPos := r.unmarshalPos(item.Position())
			item.Move(r.marshalPos(acc, crossPos))
			acc += mainSize
		}
	case justifyCenter:
		acc := remainMainAxisSizes / 2
		for _, item := range r.container.Items {
			mainSize, _ := r.unmarshalSize(item.Size())
			_, crossPos := r.unmarshalPos(item.Position())
			item.Move(r.marshalPos(acc, crossPos))
			acc += mainSize
		}
	}

	switch r.container.alignment {
	case alignmentFlexStart:
	case alignmentFlexEnd:
		for _, item := range r.container.Items {
			_, crossSize := r.unmarshalSize(item.Size())
			mainPos, _ := r.unmarshalPos(item.Position())
			item.Move(r.marshalPos(mainPos, maxCrossSize-crossSize))
		}
	case alignmentCenter:
		for _, item := range r.container.Items {
			_, crossSize := r.unmarshalSize(item.Size())
			mainPos, _ := r.unmarshalPos(item.Position())
			item.Move(r.marshalPos(mainPos, (maxCrossSize-crossSize)/2))
		}
	case alignmentStretch:
		for _, item := range r.container.Items {
			mainSize, _ := r.unmarshalSize(item.Size())
			item.Resize(r.marshalSize(mainSize, maxCrossSize))
		}
	}
}

func (r *renderer) MinSize() fyne.Size {
	var main float32
	var cross float32

	for _, item := range r.container.Items {
		m, c := r.unmarshalSize(item.Size())
		m += m
		if c > cross {
			cross = c
		}
	}

	return r.marshalSize(main, cross)
}

func (r *renderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, len(r.container.Items))
	for i := range r.container.Items {
		objects[i] = r.container.Items[i].CanvasObject
	}
	return objects
}

func (r *renderer) Destroy() {}

func (r *renderer) Refresh() {}
