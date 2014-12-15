package textbox

import "testing"

func TestRuneWidth(t *testing.T) {
	if RuneWidth('我') != 2 {
		t.Fatal("width")
	}
	if RuneWidth('A') != 1 {
		t.Fatal("width")
	}
	if DisplayWidth("吞zuo玻si璃") != 11 {
		t.Fatal("width")
	}
}
