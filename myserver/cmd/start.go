package cmd

import (
	"context"
	"fmt"
	"myserver/middleware"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func Start(ctx context.Context, wg sync.WaitGroup) {
	defer wg.Done()
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

	// go func() {
	// 	server := &http.Server{Addr: "728", Handler: r}
	// 	ln, _ := net.Listen("tcp", server.Addr)
	// 	if err := server.Serve(ln); err != nil && err != http.ErrServerClosed {
	// 		fmt.Println(nil, err.Error())
	// 	}
	// }()

	http.ListenAndServe(":728", r) //combine this router with http.Server, listen and server http request.
}
