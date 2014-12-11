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

func New() *Box {
	box := &Box{}
	return box
}
