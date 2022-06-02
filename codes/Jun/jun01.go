package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%v,%T", bytes.NewBuffer([]byte("")), bytes.NewBuffer([]byte("")))
}
