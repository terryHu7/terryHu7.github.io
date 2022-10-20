package server

import (
	"fmt"
	"log"
	"net/http"
)

func HttpServer() {
	http.HandleFunc("/http", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header, r.Body)
	})

	if err := http.ListenAndServe(":728", nil); err != nil {
		log.Fatal("ListenAndServe err: ", err)
	}
}
