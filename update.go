package textbox

import "container/list"

var geometryUpdateSequence, fillSequence *list.List

func init() {
	geometryUpdateSequence = list.New()
	fillSequence = list.New()
	initCallbacks = append(initCallbacks, func() {
		Window.geometryUpdateElement = geometryUpdateSequence.PushBack(Window)
		Window.fillElement = fillSequence.PushBack(Window)
	})
}
