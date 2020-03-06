package commonUtils

import (
	"os"
	"strings"
)

func IsTesting() bool {
	var args string = strings.Join(os.Args, " ")
	return strings.Index(args, "-test.timeout") != -1
}

// func IsTesting() bool { // this is only suitable for go 1.12
// 	os.Args
// 	return flag.Lookup("test.v") != nil
// }
