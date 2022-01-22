package flynex

import (
	fyne "fyne.io/fyne/v2"
)

var _ fyne.CanvasObject = (*FlexItem)(nil)

type FlexItem struct {
	fyne.CanvasObject
	flex int
}

func Flex(object fyne.CanvasObject, flex int) *FlexItem {
	item := &FlexItem{
		CanvasObject: object,
		flex:         flex,
	}
	return item
}

func Item(object fyne.CanvasObject) *FlexItem {
	item := &FlexItem{
		CanvasObject: object,
	}
	return item
}
