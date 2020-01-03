package commonUtils

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

type funcBeforeTerminate func()

func Terminate(msg string) {
	go func() {
		log.Println("terminating", msg)
		syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	}()
}

func CatchTerminationSignal(f funcBeforeTerminate) {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(
		sigs,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		sig := <-sigs
		log.Println("received signal:", sig)
		// v.saveVars()
		// log.Println("variables saved, ready to terminate.")
		f()
		done <- true
	}()

	log.Println("graceful termination is in place.")
	<-done
	log.Println("gracefully exiting.")
	os.Exit(123)
}
