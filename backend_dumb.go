package textbox

type DumbBackend struct {
	events chan Event
}

func NewDumbBackend() *DumbBackend {
	b := &DumbBackend{
		events: make(chan Event),
	}
	return b
}

func (b *DumbBackend) Close() {
}

func (b *DumbBackend) EventsChan() chan Event {
	return b.events
}

func (b *DumbBackend) Flush(buf *Buffer) {
}

func (b *DumbBackend) Size() (int, int) {
	return 80, 25
}
