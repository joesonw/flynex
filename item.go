package flynex

import (
	fyne "fyne.io/fyne/v2"
)

var _ fyne.CanvasObject = (*FlexItem)(nil)

type FlexItem struct {
	fyne.CanvasObject
	flex int

	mainSize  float32
	crossSize float32
}

func Flex(object fyne.CanvasObject, flex int) *FlexItem {
	item := &FlexItem{
		CanvasObject: object,
		flex:         flex,
	}
	return item
}

func MainSize(object fyne.CanvasObject, size float32) *FlexItem {
	item := &FlexItem{
		CanvasObject: object,
		mainSize:     size,
	}
	return item
}

func Item(object fyne.CanvasObject) *FlexItem {
	item := &FlexItem{
		CanvasObject: object,
	}
	return item
}
