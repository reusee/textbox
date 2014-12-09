package textbox

import "github.com/nsf/termbox-go"

var Window *Box

func init() {
	Window = New()
}

func updateWindowGeometry() {
	Window.W, Window.H = termbox.Size()
	//lg("window size %d %d\n", Window.W, Window.H)
}
