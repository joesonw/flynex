package flynex

import (
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = (*Container)(nil)
var _ fyne.WidgetRenderer = (*renderer)(nil)

const (
	alignmentStretch   = "stretch"
	alignmentFlexStart = "flex-start"
	alignmentFlexEnd   = "flex-end"
	alignmentCenter    = "center"

	justifyFlexStart    = "flex-start"
	justifyFlexEnd      = "flex-end"
	justifyCenter       = "flex-center"
	justifySpaceBetween = "space-between"
	justifySpaceEvenly  = "space-evenly"
)

type Container struct {
	widget.BaseWidget

	Items     []*FlexItem
	column    bool
	grow      bool
	alignment string
	justify   string
	debug     bool
}

func NewRowContainer(items ...*FlexItem) *Container {
	return newContainer(false, items)
}

func NewColumnContainer(items ...*FlexItem) *Container {
	return newContainer(true, items)
}

func newContainer(column bool, items []*FlexItem) *Container {
	c := &Container{
		alignment: alignmentFlexStart,
		justify:   justifyFlexStart,
		column:    column,
		Items:     items,
	}
	c.BaseWidget.ExtendBaseWidget(c)
	return c
}

func (c *Container) Debug() *Container {
	c.debug = true
	return c
}

func (c *Container) Add(item *FlexItem) *Container {
	c.Items = append(c.Items, item)
	c.Refresh()
	return c
}

func (c *Container) Remove(item *FlexItem) *Container {
	for index := range c.Items {
		if c.Items[index] == item {
			return c.RemoveAt(index)
		}
	}
	return c
}

func (c *Container) RemoveAt(index int) *Container {
	if len(c.Items) > index {
		c.Items = append(c.Items[:index], c.Items[index+1:]...)
		c.Refresh()
	}
	return c
}

func (c *Container) CreateRenderer() fyne.WidgetRenderer {
	c.BaseWidget.ExtendBaseWidget(c)
	r := &renderer{
		container: c,
	}
	return r
}

func (c *Container) Grow(grow bool) *Container {
	c.grow = grow
	return c
}

func (c *Container) AlignStretch() *Container {
	c.alignment = alignmentStretch
	return c
}

func (c *Container) AlignFlexStart() *Container {
	c.alignment = alignmentFlexStart
	return c
}

func (c *Container) AlignFlexEnd() *Container {
	c.alignment = alignmentFlexEnd
	return c
}

func (c *Container) AlignCenter() *Container {
	c.alignment = alignmentCenter
	return c
}

func (c *Container) JustifyFlexStart() *Container {
	c.justify = justifyFlexStart
	return c
}

func (c *Container) JustifyFlexEnd() *Container {
	c.justify = justifyFlexEnd
	return c
}

func (c *Container) JustifyCenter() *Container {
	c.justify = justifyCenter
	return c
}

func (c *Container) JustifySpaceBetween() *Container {
	c.justify = justifySpaceBetween
	return c
}

func (c *Container) JustifySpaceEvenly() *Container {
	c.justify = justifySpaceEvenly
	return c
}
