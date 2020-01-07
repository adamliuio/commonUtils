package commonUtils

import (
	"fmt"
	"log"
	"runtime"
)

func LogError(err error) (b bool) {
	if err != nil {
		var pc uintptr
		var line int
		pc, _, line, b = runtime.Caller(1) // using 1 to log where the error happened
		n := runtime.FuncForPC(pc).Name()
		var msg = fmt.Sprintf("[ERROR] (%s:%d)\n%+v\n", n, line, err)
		log.Printf(msg)
		// NotifyMaster(msg)
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
	log.Printf(msg)
	b = true
	return
}
