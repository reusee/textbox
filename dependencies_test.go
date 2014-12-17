package textbox

import "testing"

func TestDependencies(t *testing.T) {
	win := New(NewDumbBackend())
	defer win.Close()

	deps := NewDependencies()
	b1 := new(Box)
	b2 := new(Box)
	b3 := new(Box)
	deps.Add(b1, win.Root)
	deps.Add(b2, b1)
	deps.Add(b3, b2)
	deps.Add(b3, b1)
	deps.Add(b1, b3)

	expected := []*Box{win.Root, b1, b2, b3}
	n := 0
	cb := func(box *Box) {
		if box != expected[n] {
			t.Fatalf("%d wrong", n)
		}
		n++
	}
	deps.Iter(win.Root, cb)
}
