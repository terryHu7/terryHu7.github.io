package cmd

import (
	"context"
	"fmt"
	"myserver/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start(ctx context.Context) {
	r := gin.Default()

	//middleware
	r.Use(middleware.M1())

	r.GET("/", func(ctx *gin.Context) {
		req, _ := ctx.Get("key1") //same ctx as which in middleware??
		fmt.Println("main")
		ctx.JSON(http.StatusOK, map[string]interface{}{"req": req})
	})

	r.NoRoute(func(ctx *gin.Context) { //404
		ctx.String(http.StatusNotFound, "404 NOT FOUND, NO ROUTE")
	})
	r.Run(":728")
}
