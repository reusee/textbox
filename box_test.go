package textbox

import "testing"

func TestBox(t *testing.T) {
	topBar := New(
		func() Point { return Window.TopLeft() },
		func() Point { return Window.TopRight() },
		func(box *Box) {
		})
	bottomBar := New(
		Window.BottomLeftFunc(),
		Window.BottomRightFunc(),
		func(box *Box) {
		})
	var text string
	centerText := New(
		func() Point { return Window.Center().Move(-DisplayWidth(text)/2, 0) },
		func() Point { return Window.Center().Move(DisplayWidth(text)/2, 0) },
		func(box *Box) {
		})
	_ = topBar
	_ = bottomBar
	_ = centerText
}
