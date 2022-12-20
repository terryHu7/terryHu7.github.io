package gie

import (
	"fmt"
	"net/http"
)

type UniHandler struct {
	Router *Router
}

type HandlerFunc func(*Context)

func (uniHandler *UniHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
	if handler, ok := uniHandler.Router.Url2Handler[c.Path]; ok {
		handler(c)
	} else {
		c.Writer.Header().Set("Content-Type", "text/html")
		c.StatusCode = http.StatusNotFound
		c.Writer.WriteHeader(http.StatusNotFound)
		c.Writer.Write([]byte(fmt.Sprintf("404: %s\n", c.Path)))
	}
}
