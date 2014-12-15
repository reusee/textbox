package textbox

import (
	"sync"

	"github.com/nsf/termbox-go"
)

var (
	closed    chan struct{}
	closeOnce sync.Once

	initCallbacks []func()
)

func init() {
	closed = make(chan struct{})
}

func Init() error {
	termbox.Init()
	events := make(chan *termbox.Event)
	go func() {
		for {
			ev := termbox.PollEvent()
			switch ev.Type {
			case termbox.EventError: // NOCOVER
				Close()
				return
			default: // TODO cover this NOCOVER
				events <- &ev
			}
		}
	}()
	go func() {
		select {
		case ev := <-events: // TODO cover this NOCOVER
			switch ev.Type {
			case termbox.EventResize: // NOCOVER
				termbox.Flush()
				updateWindowGeometry()
			}
		case <-closed:
			return
		}
	}()
	for _, cb := range initCallbacks {
		cb()
	}
	return nil
}

func Close() {
	closeOnce.Do(func() {
		close(closed)
		termbox.Close()
		initCallbacks = initCallbacks[0:0]
	})
}
