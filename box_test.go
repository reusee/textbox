package textbox

import "testing"

func TestBoxBasics(t *testing.T) {
	win := New(NewDumbBackend())
	defer win.Close()

	topBar := win.Box()
	topBar.SetAdjust(
		func() Point { return win.Root.TopLeft() },
		func() Point { return win.Root.TopRight() },
		win.Root)
	topBar.SetFill(func(box *Box) {
	}, win.Root)

	bottomBar := win.Box()
	bottomBar.SetAdjust(
		win.Root.BottomLeft,
		win.Root.BottomRight,
		win.Root)
	bottomBar.SetFill(func(box *Box) {
	}, win.Root)

	centerText := win.Box()
	var text string
	centerText.SetAdjust(
		func() Point { return win.Root.Center().Move(-DisplayWidth(text)/2, 0) },
		func() Point { return win.Root.Center().Move(DisplayWidth(text)/2, 0) },
		win.Root)
	centerText.SetFill(func(box *Box) {
	}, win.Root)
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

func TestResetAdjust(t *testing.T) {
	win := New(NewDumbBackend())
	defer win.Close()

	b1 := win.Box()
	b1.SetAdjust(win.Root.TopLeft, MakeMove(win.Root.TopLeft, 1, 1), win.Root)
	b2 := win.Box()
	b2.SetAdjust(MakeMove(b1.BottomRight, 1, 1), MakeMove(b2.TopLeft, 1, 1), b1)

	if p := b1.TopLeft(); p.X != 0 || p.Y != 0 {
		t.Fatal("point")
	}
	if p := b1.BottomRight(); p.X != 1 || p.Y != 1 {
		t.Fatal("point")
	}
	if p := b2.TopLeft(); p.X != 2 || p.Y != 2 {
		t.Fatal("point")
	}
	if p := b2.BottomRight(); p.X != 3 || p.Y != 3 {
		t.Fatal("point")
	}

	b1.SetAdjust(MakeMove(win.Root.TopLeft, 3, 4), MakeMove(b1.TopLeft, 5, 6), win.Root)
	if p := b1.TopLeft(); p.X != 3 || p.Y != 4 {
		t.Fatal("point")
	}
	if p := b1.BottomRight(); p.X != 8 || p.Y != 10 {
		t.Fatal("point")
	}
	if p := b2.TopLeft(); p.X != 9 || p.Y != 11 {
		t.Fatal("point")
	}
	if p := b2.BottomRight(); p.X != 10 || p.Y != 12 {
		t.Fatal("point")
	}
}
