package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "req = %v\n", r.Body)
	})
	http.ListenAndServe(":9999", nil)
}
