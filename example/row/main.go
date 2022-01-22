package main

import (
	"image/color"

	fyne "fyne.io/fyne/v2"
	fyneapp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/joesonw/flynex"
)

var (
	Red   = color.RGBA{R: 255, A: 255}
	Green = color.RGBA{G: 255, A: 255}
	Blue  = color.RGBA{B: 255, A: 255}
)

func Rect(s float32, c color.Color) fyne.CanvasObject {
	r := canvas.NewRectangle(c)
	r.SetMinSize(fyne.NewSize(s, s))
	return r
}

func Container(items ...*flynex.FlexItem) *flynex.Container {
	c := flynex.NewRowContainer(items...)
	c.Resize(fyne.NewSize(800, 100))
	return c
}

func main() {
	app := fyneapp.New()
	root := app.NewWindow("Debugger")
	root.Resize(fyne.NewSize(800, 640))

	c := container.NewVBox()

	c.Add(widget.NewLabel("50, 60, 70, justify: flex-start"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifyFlexStart())
	c.Add(widget.NewLabel("50, 60, 70, justify: flex-end"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifyFlexEnd())
	c.Add(widget.NewLabel("50, 60, 70, justify: center"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifyCenter())
	c.Add(widget.NewLabel("50, 60, 70, justify: space-between"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifySpaceBetween())
	c.Add(widget.NewLabel("50, 60, 70, justify: space-evenly"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifySpaceEvenly())

	c.Add(widget.NewLabel("f2, f1, 70, align: flex-start"))
	c.Add(Container(
		flynex.Flex(Rect(50, Red), 2),
		flynex.Flex(Rect(60, Green), 1),
		flynex.Item(Rect(70, Blue)),
	).AlignFlexStart())
	c.Add(widget.NewLabel("f2, f1, 70, align: flex-end"))
	c.Add(Container(
		flynex.Flex(Rect(50, Red), 2),
		flynex.Flex(Rect(60, Green), 1),
		flynex.Item(Rect(70, Blue)),
	).AlignFlexEnd())
	c.Add(widget.NewLabel("f2, f1, 70, align: center"))
	c.Add(Container(
		flynex.Flex(Rect(50, Red), 2),
		flynex.Flex(Rect(60, Green), 1),
		flynex.Item(Rect(70, Blue)),
	).AlignCenter())
	c.Add(widget.NewLabel("f2, f1, 70, align: stretch"))
	c.Add(Container(
		flynex.Flex(Rect(50, Red), 2),
		flynex.Flex(Rect(60, Green), 1),
		flynex.Item(Rect(70, Blue)),
	).AlignStretch())

	c.Add(widget.NewLabel("50, 60, 70, justify: flex-start, align: flex-end"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifyFlexStart().AlignFlexEnd())
	c.Add(widget.NewLabel("50, 60, 70, justify: flex-start, align: center"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifyFlexStart().AlignCenter())
	c.Add(widget.NewLabel("50, 60, 70, justify: flex-start, align: stretch"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifyFlexStart().AlignStretch())

	c.Add(widget.NewLabel("50, 60, 70, justify: center, align: flex-end"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifyCenter().AlignFlexEnd())
	c.Add(widget.NewLabel("50, 60, 70, justify: center, align: center"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifyCenter().AlignCenter())
	c.Add(widget.NewLabel("50, 60, 70, justify: center, align: stretch"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifyCenter().AlignStretch())

	c.Add(widget.NewLabel("50, 60, 70, justify: flex-end, align: flex-end"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifyFlexEnd().AlignFlexEnd())
	c.Add(widget.NewLabel("50, 60, 70, justify: flex-end, align: center"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifyFlexEnd().AlignCenter())
	c.Add(widget.NewLabel("50, 60, 70, justify: flex-end, align: stretch"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifyFlexEnd().AlignStretch())

	c.Add(widget.NewLabel("50, 60, 70, justify: space-between, align: flex-end"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifySpaceBetween().AlignFlexEnd())
	c.Add(widget.NewLabel("50, 60, 70, justify: space-between, align: center"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifySpaceBetween().AlignCenter())
	c.Add(widget.NewLabel("50, 60, 70, justify: space-between, align: stretch"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifySpaceBetween().AlignStretch())

	c.Add(widget.NewLabel("50, 60, 70, justify: space-evenly, align: flex-end"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifySpaceEvenly().AlignFlexEnd())
	c.Add(widget.NewLabel("50, 60, 70, justify: space-evenly, align: center"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifySpaceEvenly().AlignCenter())
	c.Add(widget.NewLabel("50, 60, 70, justify: space-evenly, align: stretch"))
	c.Add(Container(
		flynex.Item(Rect(50, Red)),
		flynex.Item(Rect(60, Green)),
		flynex.Item(Rect(70, Blue)),
	).JustifySpaceEvenly().AlignStretch())

	root.SetContent(container.NewMax(container.NewVScroll(c)))
	root.Show()
	app.Run()
}
