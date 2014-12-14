package textbox

import "testing"

func TestBoxListAppend(t *testing.T) {
	l := NewBoxList()
	b1 := New(Window.TopLeft, Window.TopRight, nil)
	b2 := New(Window.TopLeft, Window.TopRight, nil)
	b3 := New(Window.TopLeft, Window.TopRight, nil)
	l.Append(b1, b2, b3)
	if l.l.Len() != 3 || len(l.m) != 3 {
		t.Fatal("size")
	}

	e := l.l.Front()
	if e.Value.(*Box) != b1 {
		t.Fatal("elem")
	}
	if l.m[b1] != e {
		t.Fatal("elem")
	}
	e = e.Next()
	if e.Value.(*Box) != b2 {
		t.Fatal("elem")
	}
	if l.m[b2] != e {
		t.Fatal("elem")
	}
	e = e.Next()
	if e.Value.(*Box) != b3 {
		t.Fatal("elem")
	}
	if l.m[b3] != e {
		t.Fatal("elem")
	}

	l.InsertAfter(b1, b3)
	e = l.l.Front()
	if e.Value.(*Box) != b1 {
		t.Fatal("elem")
	}
	if l.m[b1] != e {
		t.Fatal("elem")
	}
	e = e.Next()
	if e.Value.(*Box) != b3 {
		t.Fatal("elem")
	}
	if l.m[b3] != e {
		t.Fatal("elem")
	}
	e = e.Next()
	if e.Value.(*Box) != b2 {
		t.Fatal("elem")
	}
	if l.m[b2] != e {
		t.Fatal("elem")
	}

	l.InsertBefore(b1, b2)
	e = l.l.Front()
	if e.Value.(*Box) != b2 {
		t.Fatal("elem")
	}
	if l.m[b2] != e {
		t.Fatal("elem")
	}
	e = e.Next()
	if e.Value.(*Box) != b1 {
		t.Fatal("elem")
	}
	if l.m[b1] != e {
		t.Fatal("elem")
	}
	e = e.Next()
	if e.Value.(*Box) != b3 {
		t.Fatal("elem")
	}
	if l.m[b3] != e {
		t.Fatal("elem")
	}

}
