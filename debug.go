package textbox

import (
	"net/http"
	_ "net/http/pprof"
)

func init() {
	go http.ListenAndServe("localhost:6789", nil)
}
