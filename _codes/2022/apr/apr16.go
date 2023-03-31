package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 0
	// for {
	// 	_, err := fmt.Scanf("%d%d", &a, &b)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			fmt.Print("aaa")
	// 			break
	// 		}
	// 	} else {
	// 		fmt.Print("eee")
	// 		fmt.Printf("%d\n", a+b)
	// 	}
	// }
	for {
		_, err := fmt.Scanf("%d%d", &a, &b)
		for err == nil {
			fmt.Print(a, b)
		}
	}

}
