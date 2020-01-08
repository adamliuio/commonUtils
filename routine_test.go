package commonUtils

import (
	"log"
	"testing"
	"time"
)

func one() {
	log.Println("one")
}

func two() {
	log.Println("two")
}

func TestRoutine(t *testing.T) {
	defer wg.Wait()
	funcs := []FuncToLoop{one, two}
	slps := []time.Duration{time.Second, time.Second * 2}
	GoroutineFuncs(funcs, slps)
}
