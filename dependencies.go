package textbox

import "sync"

var adjustDependencies, fillDependencies *Dependencies

func init() {
	adjustDependencies = NewDependencies()
	fillDependencies = NewDependencies()
}

type Dependencies struct {
	sync.RWMutex
	m map[*Box]map[*Box]struct{}
}

func NewDependencies() *Dependencies {
	return &Dependencies{
		m: make(map[*Box]map[*Box]struct{}),
	}
}

func (d *Dependencies) Add(box, depend *Box) {
	d.Lock()
	defer d.Unlock()
	m, ok := d.m[depend]
	if !ok {
		m = make(map[*Box]struct{})
		d.m[depend] = m
	}
	m[box] = struct{}{}
}

func (d *Dependencies) Iter(from *Box, fn func(*Box)) {
	itered := make(map[*Box]struct{})
	d.iter(from, fn, itered)
}

func (d *Dependencies) iter(from *Box, fn func(*Box), itered map[*Box]struct{}) {
	if _, ok := itered[from]; ok {
		return
	}
	fn(from)
	itered[from] = struct{}{}
	for next := range d.m[from] {
		d.iter(next, fn, itered)
	}
}
