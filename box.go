package textbox

type Point struct {
	X, Y int
}

func (p Point) Move(x, y int) Point {
	return Point{p.X + x, p.Y + y}
}

type Box struct {
	win                          *Window
	topLeft, bottomRight         Point
	topLeftFunc, bottomRightFunc func() Point
	fillFunc                     func(*Box)
}

func (w *Window) Box() *Box {
	return &Box{
		win: w,
	}
}

func (b *Box) SetAdjust(topLeft, bottomRight func() Point, depends ...*Box) {
	b.topLeftFunc = topLeft
	b.bottomRightFunc = bottomRight
	for _, box := range depends {
		b.win.adjustDependencies.Add(b, box)
	}
	b.adjust()
}

func (b *Box) adjust() {
	b.win.adjustDependencies.Iter(b, func(box *Box) {
		box.topLeft = box.topLeftFunc()
		box.bottomRight = box.bottomRightFunc()
	})
}

func (b *Box) SetFill(fill func(*Box), depends ...*Box) {
	b.fillFunc = fill
	for _, box := range depends {
		b.win.fillDependencies.Add(b, box)
	}
	b.fill()
}

func (b *Box) fill() {
	//TODO
}

func (b *Box) TopLeft() Point {
	return b.topLeft
}
func (b *Box) TopCenter() Point {
	return Point{(b.topLeft.X + b.bottomRight.X) / 2, b.topLeft.Y}
}
func (b *Box) TopRight() Point {
	return Point{b.bottomRight.X, b.topLeft.Y}
}
func (b *Box) CenterLeft() Point {
	return Point{b.topLeft.X, (b.topLeft.Y + b.bottomRight.Y) / 2}
}
func (b *Box) Center() Point {
	return Point{(b.topLeft.X + b.bottomRight.X) / 2, (b.topLeft.Y + b.bottomRight.Y) / 2}
}
func (b *Box) CenterRight() Point {
	return Point{b.bottomRight.X, (b.topLeft.Y + b.bottomRight.Y) / 2}
}
func (b *Box) BottomLeft() Point {
	return Point{b.topLeft.X, b.bottomRight.Y}
}
func (b *Box) BottomCenter() Point {
	return Point{(b.topLeft.X + b.bottomRight.X) / 2, b.bottomRight.Y}
}
func (b *Box) BottomRight() Point {
	return b.bottomRight
}
