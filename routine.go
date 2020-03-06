package commonUtils

import (
	"sync"
	"time"
)

type FuncToLoop func()
type FuncSchedule struct {
	Function FuncToLoop
	SleepFor time.Duration
}

var wg sync.WaitGroup = sync.WaitGroup{}

func GoroutineFuncs(funcSchedules []FuncSchedule) {
	for _, fs := range funcSchedules {
		go funcWithForLoop(fs.Function, fs.SleepFor)()
	}
}

func funcWithForLoop(f FuncToLoop, duration time.Duration) (ff FuncToLoop) {
	wg.Add(1)
	return func() {
		for {
			f()
			time.Sleep(duration)
		}
	}
}
