package commonUtils

import (
	"log"
	"path/filepath"
	"runtime"
)

func GetDirFileLine() (dir, file string, line int) {
	var pc uintptr
	pc, file, line, _ = runtime.Caller(1) // using 1 to log where the error happened
	dir, file = filepath.Split(file)
	return
}
