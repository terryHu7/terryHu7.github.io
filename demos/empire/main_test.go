package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMain(t *testing.T) {
	http.HandleFunc("/", samplehandler)
	http.ListenAndServe(":9999", nil)
}

func samplehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("w:%+v\n", w)
	fmt.Printf("r:%+v\n", r)
	fmt.Fprintf(w, "req = %v\n", r.Body)
}
