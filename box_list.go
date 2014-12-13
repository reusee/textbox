package textbox

import "container/list"

type BoxList struct {
	l *list.List
	m map[*Box]*list.Element
}

func NewBoxList() *BoxList {
	return &BoxList{
		l: list.New(),
		m: make(map[*Box]*list.Element),
	}
}
