package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := "hello"
	h := md5.New()
	h.Write([]byte(data))
	fmt.Printf("%x", h.Sum(nil))
}
