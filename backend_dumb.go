package textbox

type DumbBackend struct {
	width, height int
	events        chan Event
}

func NewDumbBackend() *DumbBackend {
	b := &DumbBackend{
		width:  80,
		height: 25,
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
	return b.width, b.height
}

func (b *DumbBackend) Resize(width, height int) {
	b.width = width
	b.height = height
	b.events <- ResizeEvent{width, height}
}
