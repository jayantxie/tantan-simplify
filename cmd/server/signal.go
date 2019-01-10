package main

import (
	"os"
	"os/signal"
	"syscall"
)

func handleSignals(sig os.Signal, logsReopenCallback func()) (exitNow bool) {
	switch sig {
	case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
		return true
	case syscall.SIGUSR1:
		logsReopenCallback()
		return false
	}
	return false
}

func registerSignal(shutdown chan struct{}, logsReopenCallback func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM}...)
	go func() {
		for sig := range c {
			if handleSignals(sig, logsReopenCallback) {
				close(shutdown)
				return
			}
		}
	}()
}
