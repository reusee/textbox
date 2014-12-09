package textbox

import (
	"fmt"
	"log"
	"os"
)

var (
	sp     = fmt.Sprintf
	pt     = fmt.Printf
	logger = log.New(os.Stdout, "", 0)
	lg     = logger.Printf
)
