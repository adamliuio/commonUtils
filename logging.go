package commonUtils

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func LogError(err error) (b bool) {
	if err != nil {
		var pc uintptr
		var file string
		var line int
		pc, file, line, b = runtime.Caller(1) // using 1 to log where the error happened
		file = filepath.Base(file)
		n := runtime.FuncForPC(pc).Name()
		var msg = fmt.Sprintf("[ERROR] %s (%s:%d)\n%+v\n", n, file, line, err)
		fmt.Printf(msg)
		b = true
	}
	return
}

func LogString(s string) (b bool) {
	var pc uintptr
	var line int
	pc, _, line, b = runtime.Caller(1)
	n := runtime.FuncForPC(pc).Name()
	var msg = fmt.Sprintf("[ERROR] (%s:%d)\n%s\n", n, line, s)
	fmt.Printf(msg)
	b = true
	return
}
