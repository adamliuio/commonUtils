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

func TestRunEveryDayAt(t *testing.T) {
	RunEveryDayAt(damn, 20, 06, 40, 0)
}
