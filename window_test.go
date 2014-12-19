package textbox

import "testing"

func TestResize(t *testing.T) {
	backend := NewDumbBackend()
	win := New(backend)

	width := 40
	height := 30
	backend.Resize(width, height)
	if w, h := win.backend.Size(); w != width || h != height {
		t.Fatal("size")
	}
	//TODO wait for event?
	if len(win.buffer.bs) != width*height {
		t.Fatal("buffer length")
	}

	width = 40
	height = 30
	backend.Resize(width, height)
	if w, h := win.backend.Size(); w != width || h != height {
		t.Fatal("size")
	}
	if len(win.buffer.bs) != width*height {
		t.Fatal("buffer length")
	}
}
