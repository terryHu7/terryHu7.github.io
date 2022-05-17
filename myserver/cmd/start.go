package cmd

import (
	"context"
	"fmt"
	"myserver/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start(ctx context.Context) {

	// defer wg.Done()
	r := gin.Default()

	//middleware
	r.Use(middleware.M1())

	r.GET("/test", func(ctx *gin.Context) {
		req, _ := ctx.Get("key1") //same ctx as which in middleware??
		fmt.Println("main")
		ctx.JSON(http.StatusOK, map[string]interface{}{"req": req})
	})

	r.NoRoute(func(ctx *gin.Context) { //404
		ctx.String(http.StatusNotFound, "404 NOT FOUND, NO ROUTE")
	})

	// r.Run(":728") //gin

	server := &http.Server{Addr: ":728", Handler: r} //combine this router with http.Server
	server.ListenAndServe()                          // listen and server http request
	<-ctx.Done()                                     //release port
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: " + err.Error())
	}

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
