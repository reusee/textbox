package textbox

import (
	"testing"
	"time"
)

func TestDumbBackend(t *testing.T) {
	backend := NewDumbBackend()
	win := New(backend)
	defer win.Close()

	go func() {
		for {
			ev, ok := <-win.Events
			if !ok {
				return
			}
			_ = ev
		}
	}()

	time.Sleep(time.Second)
	backend.events <- ResizeEvent{800, 600}
	backend.events <- ErrorEvent{}
}
