package main

import "fmt"

func main() {
	// := work in the code block
	a := 1
	if a := 2; !true {
		fmt.Println(a)
	} else {
		a := 3
		fmt.Println(a)
	}
	fmt.Println(a)
}
