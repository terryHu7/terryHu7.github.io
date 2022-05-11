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

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		cancel()
	}()

	cmd.Start(ctx)
}
