package handlers

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func InterruptHandler(errc chan<- error) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	terminateError := fmt.Errorf("%s", <-c)

	fmt.Println("service shutdown...")
	// Place whatever shutdown handling you want here

	errc <- terminateError
}
