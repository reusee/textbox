package textbox

type Event interface {
	isEvent()
}

type ErrorEvent struct {
}

func (e ErrorEvent) isEvent() {} // NOCOVER

type ResizeEvent struct {
	Width, Height int
}

func (e ResizeEvent) isEvent() {} // NOCOVER
