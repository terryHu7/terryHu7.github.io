package main

import (
	"context"
	"myserver/cmd"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, os.Interrupt)
		if <-c != nil {
			cancel()
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	cmd.Start(ctx, wg)
	wg.Wait()
}
