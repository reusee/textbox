package textbox

import (
	"sync"
)

type Window struct {
	backend                              Backend
	closed                               chan struct{}
	closeOnce                            sync.Once
	buffer                               *Buffer
	adjustDependencies, fillDependencies *Dependencies

	Events chan Event
	Root   *Box
}

func New(backend Backend) *Window {
	width, height := backend.Size()
	win := &Window{
		backend:            backend,
		closed:             make(chan struct{}),
		buffer:             NewBuffer(width * height),
		adjustDependencies: NewDependencies(),
		fillDependencies:   NewDependencies(),

		Events: make(chan Event),
	}
	win.Root = win.Box()
	win.Root.SetAdjust(func() Point { return Point{0, 0} },
		func() Point {
			w, h := backend.Size()
			return Point{w, h}
		})
	win.Root.Adjust()
	eventsChan := backend.EventsChan()
	go func() {
		for {
			select {
			case ev := <-eventsChan:
				switch ev := ev.(type) {
				case ErrorEvent:
					win.Close()
					return
				case ResizeEvent:
					win.buffer.Resize(ev.Width * ev.Height)
					win.Root.Adjust()
					win.Root.Fill()
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

	return win
}

func (w *Window) Close() {
	w.closeOnce.Do(func() {
		close(w.closed)
		w.backend.Close()
		close(w.Events)
	})
}
