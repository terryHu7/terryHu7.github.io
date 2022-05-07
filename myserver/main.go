package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//middleware
	r.Use(m1())

	r.GET("/", func(c *gin.Context) {
		req, _ := c.Get("key1") //same ctx as which in middleware??
		fmt.Println("main")
		c.JSON(http.StatusOK, map[string]interface{}{"req": req})
	})

	r.Run(":728")
}
