package textbox

import (
	"sync"
)

var (
	backend   Backend
	closed    chan struct{}
	closeOnce sync.Once
	buffer    *Buffer

	initCallbacks []func()
	Events        chan Event
)

func Init(backend Backend) error {
	backend = backend
	closed = make(chan struct{})
	buffer = NewBuffer()
	Events = make(chan Event)

	eventsChan := backend.EventsChan()
	go func() {
		for {
			select {
			case ev := <-eventsChan:
				switch ev.(type) {
				case ErrorEvent:
					Close()
					return
				case ResizeEvent:
					updateWindowGeometry()
					backend.Flush(buffer)
				}
				select {
				case Events <- ev:
				default:
				}
			case <-closed:
				return
			}
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
		backend.Close()
		initCallbacks = initCallbacks[0:0]
		close(Events)
	})
}
