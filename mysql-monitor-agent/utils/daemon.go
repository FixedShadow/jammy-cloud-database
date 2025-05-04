package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func StartDaemon() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	<-c
	close(c)
	os.Exit(0)
}
