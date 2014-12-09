package textbox

import "container/list"

var layoutSequence, fillSequence *list.List

func init() {
	layoutSequence = list.New()
	fillSequence = list.New()
	initCallbacks = append(initCallbacks, func() {
		Window.layoutElement = layoutSequence.PushBack(Window)
		Window.fillElement = fillSequence.PushBack(Window)
	})
}
