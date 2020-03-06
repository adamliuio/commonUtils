package commonUtils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type funcBeforeTerminate func()

func Terminate(msg string) {
	go func() {
		fmt.Println("terminating", msg)
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
		fmt.Println("received signal:", sig)
		// v.saveVars()
		// fmt.Println("variables saved, ready to terminate.")
		f()
		done <- true
	}()

	fmt.Println("graceful termination is in place.")
	<-done
	fmt.Println("gracefully exiting.")
	os.Exit(123)
}
