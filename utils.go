package textbox

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/text/width"
)

var (
	sp     = fmt.Sprintf
	pt     = fmt.Printf
	logger = log.New(os.Stdout, "", 0)
	lg     = logger.Printf
)

func RuneWidth(r rune) int {
	prop := width.LookupRune(r)
	switch prop.Kind() {
	case width.EastAsianWide, width.EastAsianAmbiguous, width.EastAsianFullwidth:
		return 2
	}
	return 1
}

func DisplayWidth(s string) (ret int) {
	for _, r := range s {
		ret += RuneWidth(r)
	}
	return
}

func MakeMove(fn func() Point, x, y int) func() Point {
	return func() Point {
		return fn().Move(x, y)
	}
}
