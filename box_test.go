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
	//TODO
}

func TestBoxPoints(t *testing.T) {
	box := &Box{
		topLeft:     Point{1, 2},
		bottomRight: Point{7, 8},
	}

	if p := box.TopLeft(); p.X != 1 || p.Y != 2 {
		t.Fatal("point")
	}
	if p := box.TopCenter(); p.X != 4 || p.Y != 2 {
		t.Fatal("point")
	}
	if p := box.TopRight(); p.X != 7 || p.Y != 2 {
		t.Fatal("point")
	}
	if p := box.CenterLeft(); p.X != 1 || p.Y != 5 {
		t.Fatal("point")
	}
	if p := box.Center(); p.X != 4 || p.Y != 5 {
		t.Fatal("point")
	}
	if p := box.CenterRight(); p.X != 7 || p.Y != 5 {
		t.Fatal("point")
	}
	if p := box.BottomLeft(); p.X != 1 || p.Y != 8 {
		t.Fatal("point")
	}
	if p := box.BottomCenter(); p.X != 4 || p.Y != 8 {
		t.Fatal("point")
	}
	if p := box.BottomRight(); p.X != 7 || p.Y != 8 {
		t.Fatal("point")
	}

	if p := box.TopLeftFunc()(); p.X != 1 || p.Y != 2 {
		t.Fatal("point")
	}
	if p := box.TopCenterFunc()(); p.X != 4 || p.Y != 2 {
		t.Fatal("point")
	}
	if p := box.TopRightFunc()(); p.X != 7 || p.Y != 2 {
		t.Fatal("point")
	}
	if p := box.CenterLeftFunc()(); p.X != 1 || p.Y != 5 {
		t.Fatal("point")
	}
	if p := box.CenterFunc()(); p.X != 4 || p.Y != 5 {
		t.Fatal("point")
	}
	if p := box.CenterRightFunc()(); p.X != 7 || p.Y != 5 {
		t.Fatal("point")
	}
	if p := box.BottomLeftFunc()(); p.X != 1 || p.Y != 8 {
		t.Fatal("point")
	}
	if p := box.BottomCenterFunc()(); p.X != 4 || p.Y != 8 {
		t.Fatal("point")
	}
	if p := box.BottomRightFunc()(); p.X != 7 || p.Y != 8 {
		t.Fatal("point")
	}
}
