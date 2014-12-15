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
	termbox.Init()
	closed = make(chan struct{})
}

func Init() error {
	events := make(chan *termbox.Event)
	go func() {
		for {
			ev := termbox.PollEvent()
			switch ev.Type {
			case termbox.EventError:
				Close()
				return
			default:
				events <- &ev
			}
		}
	}()
	go func() {
		select {
		case ev := <-events:
			switch ev.Type {
			case termbox.EventResize:
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
