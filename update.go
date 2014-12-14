package textbox

var adjustList, fillList *BoxList

func init() {
	adjustList = NewBoxList()
	fillList = NewBoxList()

	initCallbacks = append(initCallbacks, func() {
		adjustList.Append(Window)
		fillList.Append(Window)
	})
}
