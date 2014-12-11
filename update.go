package textbox

import "container/list"

var adjustSequence, fillSequence *list.List

func init() {
	adjustSequence = list.New()
	fillSequence = list.New()
	initCallbacks = append(initCallbacks, func() {
		Window.adjustElement = adjustSequence.PushBack(Window)
		Window.fillElement = fillSequence.PushBack(Window)
	})
}
