package textbox

type Buffer struct {
	bs []byte
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		bs: make([]byte, size),
	}
}

func (b *Buffer) Resize(size int) {
	if size == len(b.bs) {
		return
	}
	if size > len(b.bs) {
		b.bs = append(b.bs, make([]byte, size-len(b.bs))...)
	} else {
		b.bs = b.bs[:size]
	}
}
