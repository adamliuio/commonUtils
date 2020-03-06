package commonUtils

import (
	"log"
	"testing"
	"time"
)

func damn() {
	log.Println(time.Now())
	log.Println("do some job.")
}

func TestRunEveryAt(t *testing.T) {
	go RunEveryDayAt(damn, 13, 10, 35, 0)
	go RunEveryHourAt(damn, 10, 35, 0)
	RunEveryMinuteAt(damn, 35, 0)
}

func _TestRunEvery(t *testing.T) { // passed
	RunEvery(damn, 2, "Second", true)
}

func _TestRunEveryNSecond(t *testing.T) {
	RunEveryNSecond(damn, 5, true)
}

func _TestRunEveryDayAt(t *testing.T) {
	RunEveryDayAt(damn, 12, 14, 50, 0)
}
