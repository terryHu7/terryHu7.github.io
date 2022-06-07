package main

import (
	"fmt"
	"io"
)

func main() {

	arr := []int{}
	n := 0
	lines := 2
	for {
		for i := 0; i < lines; i++ {
			_, err1 := fmt.Scanln(&arr)
			if err1 != nil {
				if err1 == io.EOF {
					break
				}
			}
			_, err2 := fmt.Scanln(&n)
			if err2 != nil {
				if err2 == io.EOF {
					break
				}
			}
		}
		fmt.Println(minKBitFlips(arr, n))
	}

}

func minKBitFlips(A []int, K int) int {
	flippedTime, count := 0, 0
	for i := 0; i < len(A); i++ {
		if i >= K && A[i-K] == 2 {
			flippedTime--
		}
		if flippedTime%2 == A[i] {
			if i+K > len(A) {
				return -1
			}
			A[i] = 2
			flippedTime++
			count++
		}
	}
	return count
}
