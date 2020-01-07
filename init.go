package commonUtils

import (
	"flag"
	"log"
	"os"
)

var (
	Mode string
)

func init() {
	os.Setenv("TZ", "Asia/Shanghai")
	if flag.Lookup("test.v") == nil {
		Mode = "normal"
	} else {
		Mode = "testing"
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
