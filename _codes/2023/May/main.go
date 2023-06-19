package main

import (
	"os"
)

func main() {
	data := []byte("hhh")
	f, err := os.OpenFile("tmp.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	f.Write(data)
	defer f.Close()
	f.Sync()
}
