package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("main start")

	defer func() {
		quit()
	}()
	ctx, cancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		s := <-sig
		cancel()
		fmt.Printf("receive signal: [%v]\n", s)
		time.Sleep(time.Second)
		os.Exit(88)
	}()

	go func(ctx context.Context) {
		<-ctx.Done()
		fmt.Println("<-ctx.Done, no need to carry on")
	}(ctx)

	<-time.After(10 * time.Second)
	fmt.Println("10s time out")

	fmt.Println("main end")
}

func quit() {
	fmt.Println("\n退出")
}
