package gie

import (
	"fmt"
	"net/http"
)

func Run() {
	uniHandler := UniHandler{
		Router: &Router{
			Url2Handler: make(map[string]HandlerFunc),
		},
	}
	uniHandler.Router.Url2Handler["/test"] = func(ctx *Context) {
		fmt.Println("test")
	}
	http.ListenAndServe("0.0.0.0:2404", &uniHandler)
}
