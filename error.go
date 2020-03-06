package commonUtils

import (
	"fmt"
	"os"
	"runtime"
)

func CheckError(err error, mode, url string) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1) // using 1 to log where the error happened
		var loc string = fmt.Sprintf("%s:%d", file, line)
		var msg string = loc + ": err: " + err.Error()
		if mode == "release" {
			NotifyMaster(msg, url)
		} else if mode == "debug" {
			println(msg)
			os.Exit(1)
		} else {
			println("mode can only be 'release' or 'debug'.")
			os.Exit(1)
		}
	}
}
