package textbox

import "testing"

func TestBox(t *testing.T) {
	topBar := New()
	topBar.SetAdjust(
		func() Point { return Window.TopLeft() },
		func() Point { return Window.TopRight() },
		Window)
	topBar.SetFill(func(box *Box) {
	}, Window)

	bottomBar := New()
	bottomBar.SetAdjust(
		Window.BottomLeft,
		Window.BottomRight,
		Window)
	bottomBar.SetFill(func(box *Box) {
	}, Window)

	centerText := New()
	var text string
	centerText.SetAdjust(
		func() Point { return Window.Center().Move(-DisplayWidth(text)/2, 0) },
		func() Point { return Window.Center().Move(DisplayWidth(text)/2, 0) },
		Window)
	centerText.SetFill(func(box *Box) {
	}, Window)

	_ = topBar
	_ = bottomBar
	_ = centerText
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
}

func TestPointMove(t *testing.T) {
	if p := (Point{0, 0}).Move(1, 1); p.X != 1 || p.Y != 1 {
		t.Fatal("move")
	}
	if p := (Point{5, 6}).Move(-7, -8); p.X != -2 || p.Y != -2 {
		t.Fatal("move")
	}
}
