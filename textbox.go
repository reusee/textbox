package textbox

import "github.com/nsf/termbox-go"

func Init() error {
	if err := termbox.Init(); err != nil {
		return err
	}
	return nil
}

func Close() {
	termbox.Close()
}
