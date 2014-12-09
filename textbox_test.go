package textbox

import (
	"testing"
	"time"
)

func TestInitAndClose(t *testing.T) {
	Init()
	Close()
}

func TestWindowSize(t *testing.T) {
	Init()
	time.Sleep(time.Second)
	if Window.W == 0 || Window.H == 0 {
		t.Fatal("window size")
	}
	Close()
}
