package textbox

import (
	"testing"
	"time"
)

func TestWindowSize(t *testing.T) {
	Init()
	time.Sleep(time.Second)
	if Window.bottomRight.X == 0 || Window.bottomRight.Y == 0 {
		t.Fatal("window size")
	}
	Close()
}
