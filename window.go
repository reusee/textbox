package textbox

import "github.com/nsf/termbox-go"

var Window *Box

func init() {
	Window = &Box{}
	initCallbacks = append(initCallbacks, func() {
		updateWindowGeometry()
	})
}

func updateWindowGeometry() {
	w, h := termbox.Size()
	Window.bottomRight = Point{w, h}
}
