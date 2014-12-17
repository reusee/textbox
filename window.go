package textbox

var Window *Box

func init() {
	Window = &Box{}
	initCallbacks = append(initCallbacks, func() {
		updateWindowGeometry()
	})
}

func updateWindowGeometry() {
	w, h := backend.Size()
	Window.bottomRight = Point{w, h}
}
