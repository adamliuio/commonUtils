package commonUtils

import (
	"log"
	"testing"
	"time"
)

func one() {
	log.Println("1")
}

func two() {
	log.Println("1 + 1")
}

func TestRoutine(t *testing.T) {
	defer wg.Wait()
	GoroutineFuncs([]FuncSchedule{
		FuncSchedule{Function: one, SleepFor: time.Second},
		FuncSchedule{Function: two, SleepFor: time.Second * 2},
	})
}
