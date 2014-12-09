package textbox

import "container/list"

type Box struct {
	X, Y, W, H    int
	layoutElement *list.Element
	fillElement   *list.Element
}

func New() *Box {
	box := &Box{}
	return box
}
