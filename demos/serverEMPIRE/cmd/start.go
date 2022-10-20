package cmd

import (
	"context"
)

func Start(ctx context.Context) {

	// defer wg.Done()
	/*
			syncq, _ := net.Listen("tcp", "127.0.0.1:728")
		LOOP:
			for {
				acceptq, _ := syncq.Accept()
				select {
				case <-ctx.Done():
					acceptq.Close()
					break LOOP
				default:
					go func(acceptq net.Conn, ctx context.Context) {
						defer acceptq.Close()
						fmt.Println(acceptq.LocalAddr(), acceptq.RemoteAddr())
					}(acceptq, ctx)
				}

			}
	*/
}
