package commonUtils

import (
	"sync"
	"time"
)

type FuncToLoop func()

var wg sync.WaitGroup = sync.WaitGroup{}

func funcWithForLoop(f FuncToLoop, duration time.Duration) (ff FuncToLoop) {
	wg.Add(1)
	return func() {
		for {
			f()
			time.Sleep(duration)
		}
	}
}

func GoroutineFuncs(fs []FuncToLoop, sleepFor []time.Duration) {
	for i, f := range fs {
		go funcWithForLoop(f, sleepFor[i])()
	}
}
