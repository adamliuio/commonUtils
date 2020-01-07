package commonUtils

import (
	"log"
	"testing"
	"time"
)

func ok() {
	log.Println("ok if this is printed means the shit worked.")
}

func _TestTerminate(t *testing.T) {
	go func() {
		CatchTerminationSignal(ok)
	}()
	log.Println("right after CatchTerminationSignal.")
	time.Sleep(3 * time.Second)
	Terminate("the test process.")
	time.Sleep(3 * time.Second)
	log.Fatalln("terminate() didn't work")
}
