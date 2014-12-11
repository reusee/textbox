package textbox

import "testing"

func TestBox(t *testing.T) {
	top := New(
		Align{Window, TopLeft, 0, 0},
		Align{Window, TopRight, 0, 0},
		Fill{func() {
		}, Window})
	bottom := New(
		Align{Window, BottomLeft, 0, 0},
		Align{Window, BottomRight, 0, 0},
		Fill{func() {
		}, Window})
	middle := New(
		Align{top, BottomLeft, 0, 1},
		Align{bottom, TopRight, 0, -1},
		Fill{func() {
		}, Window})
	//TODO
	_ = top
	_ = bottom
	_ = middle
}
