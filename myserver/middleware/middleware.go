package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func M1() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := fmt.Println
		tpin := time.Now()
		ctx.Set("key1", "value1")
		p("ctx.Writer.Status(): ", ctx.Writer.Status())
		ctx.Next() //do it when all middlewares and handlers finish, which is usually used to calculate the cost time.
		// if here is `return`, the code left will be useless, but middlewares/handlers will be run .
		// if here's `ctx.Abort()`, the next middleware/handler will be useless, but the code left in this middleware will be run as well.
		costtime := time.Since(tpin)
		p("costtime: ", costtime)
	}
}
