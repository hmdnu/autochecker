package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/hmdnu/autocheck/service"
)

func main() {
	service.AutoCheck()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
