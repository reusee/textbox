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
	if Window.BottomRight.X == 0 || Window.BottomRight.Y == 0 {
		t.Fatal("window size")
	}
	Close()
}
