package main

import (
	"fmt"
	"unsafe"
)

func main() {
	b := make([]int, 4, 5)
	b1 := b[:3]
	b2 := b[1:]
	fmt.Println(b, b1, b2, unsafe.Sizeof(b), unsafe.Sizeof(&b[0]), unsafe.Sizeof("absdgsdfgsafc"))
	fmt.Printf("b: %p, b 1st elment: %p\n", b, &b[1])
	c := new(chan int)
	fmt.Println(c)
}
