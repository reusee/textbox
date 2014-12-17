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
				win.Events <- ev
			case <-win.closed:
				return
			}
		}
	}()

	win.updateWindowGeometry()

	return win
}

func (w *Window) updateWindowGeometry() {
	width, height := w.backend.Size()
	w.Root.bottomRight = Point{width, height}
}

func (w *Window) Close() {
	w.closeOnce.Do(func() {
		close(w.closed)
		w.backend.Close()
		close(w.Events)
	})
}
