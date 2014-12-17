package textbox

import (
	"sync"
)

type Window struct {
	backend   Backend
	closed    chan struct{}
	closeOnce sync.Once
	buffer    *Buffer

	Events chan Event
	Root   *Box
}

func New(backend Backend) *Window {
	win := &Window{
		backend: backend,
		closed:  make(chan struct{}),
		buffer:  NewBuffer(),

		Events: make(chan Event),
	}
	win.Root = win.Box()
	eventsChan := backend.EventsChan()
	go func() {
		for {
			select {
			case ev := <-eventsChan:
				switch ev.(type) {
				case ErrorEvent:
					win.Close()
					return
				case ResizeEvent:
					win.updateWindowGeometry()
					backend.Flush(win.buffer)
				}
				select {
				case win.Events <- ev:
				default:
				}
			case <-win.closed:
				return
			}
		}
	}()

	win.updateWindowGeometry()

	return win
}

func (t *Window) updateWindowGeometry() {
	w, h := t.backend.Size()
	t.Root.bottomRight = Point{w, h}
}

func (t *Window) Close() {
	t.closeOnce.Do(func() {
		close(t.closed)
		t.backend.Close()
		close(t.Events)
	})
}
