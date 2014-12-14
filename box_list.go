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

func (l *BoxList) Append(boxes ...*Box) {
	for _, box := range boxes {
		elem := l.l.PushBack(box)
		l.m[box] = elem
	}
}

func (l *BoxList) InsertAfter(target, box *Box) {
	elem := l.l.InsertAfter(box, l.m[target])
	l.m[box] = elem
}

func (l *BoxList) InsertBefore(target, box *Box) {
	elem := l.l.InsertBefore(box, l.m[target])
	l.m[box] = elem
}
