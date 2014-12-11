package textbox

import "container/list"

type Point struct {
	X, Y int
}

type Box struct {
	TopLeft, BottomRight Point
	adjustElement        *list.Element
	fillElement          *list.Element
}

type Align struct {
	To               *Box
	Position         Position
	XOffset, YOffset int
}

type Position int

const (
	TopLeft Position = iota
	TopRight
	BottomLeft
	BottomRight
)

type Fill struct {
	Func  func()
	Above *Box
}

func New(topLeftAlign, bottomRightAlign Align, fill Fill) *Box {
	box := &Box{}
	return box
}
