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
	Window.BottomRight.X, Window.BottomRight.Y = termbox.Size()
}
