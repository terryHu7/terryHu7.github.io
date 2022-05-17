package main

import (
	"context"
	"myserver/cmd"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	cmd.Start(ctx)

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, os.Interrupt)
	if <-c != nil {
		cancel()
	}
	<-ctx.Done()
	// var wg sync.WaitGroup
	// wg.Add(1)
	// wg.Wait()
}
