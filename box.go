package textbox

type Box struct {
	X, Y, W, H int
}

func New() *Box {
	box := &Box{}
	return box
}
