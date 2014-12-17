package textbox

type Point struct {
	X, Y int
}

func (p Point) Move(x, y int) Point {
	return Point{p.X + x, p.Y + y}
}

type Box struct {
	topLeft, bottomRight         Point
	topLeftFunc, bottomRightFunc func() Point
	fillFunc                     func(*Box)
}

func (t *Window) Box() *Box {
	return new(Box)
}

func (b *Box) SetAdjust(topLeft, bottomRight func() Point, depends ...*Box) {
	b.topLeftFunc = topLeft
	b.bottomRightFunc = bottomRight
	for _, box := range depends {
		adjustDependencies.Add(b, box)
	}
}

func (b *Box) SetFill(fill func(*Box), depends ...*Box) {
	b.fillFunc = fill
	for _, box := range depends {
		fillDependencies.Add(b, box)
	}
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
