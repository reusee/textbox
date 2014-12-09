package textbox

import "github.com/nsf/termbox-go"

var Window *Box

func init() {
	Window = New()
	initCallbacks = append(initCallbacks, func() {
		updateWindowGeometry()
	})
}

func updateWindowGeometry() {
	Window.W, Window.H = termbox.Size()
	//lg("window size %d %d\n", Window.W, Window.H)
}
