package textbox

type Backend interface {
	Size() (int, int)
	EventsChan() chan Event
	Flush(*Buffer)
	Close()
}
