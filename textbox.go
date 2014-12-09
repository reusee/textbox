package textbox

import (
	"sync"

	"github.com/nsf/termbox-go"
)

var (
	closed    chan struct{}
	closeOnce sync.Once
)

func init() {
	closed = make(chan struct{})
}

func Init() error {
	if err := termbox.Init(); err != nil {
		return err
	}
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
			_ = ev
			//TODO
		case <-closed:
			return
		}
	}()
	return nil
}

func Close() {
	closeOnce.Do(func() {
		close(closed)
		termbox.Close()
	})
}
