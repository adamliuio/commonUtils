package commonUtils

import (
	"sync"
	"time"
)

type funcToLoop func()

var wg sync.WaitGroup = sync.WaitGroup{}

func funcWithForLoop(f funcToLoop, duration time.Duration) (ff funcToLoop) {
	wg.Add(1)
	return func() {
		for {
			f()
			time.Sleep(duration)
		}
	}
}

func GoroutineFuncs(fs []funcToLoop, sleepFor []time.Duration) {
	for i, f := range fs {
		go funcWithForLoop(f, sleepFor[i])()
	}
}
