package main

import (
	"fmt"
)

func main() {
	m := [...]int{
		'a': 1,
		'b': 2,
		'c': 3,
	}
	m['a'] = 3
	fmt.Printf("m:%T,%+v\n", m, m)
	fmt.Println(m[0], len(m))
	fmt.Println([...]int{1: 4, 5})
}
