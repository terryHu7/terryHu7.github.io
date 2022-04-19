package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// arr := ""
	n := 0
	// lines := 2
	for {
		// for i := 0; i < lines; i++ {
		input := bufio.NewScanner(os.Stdin)
		arr := input.Text()
		// for _, ss := range inputStr[1 : len(inputStr)-1] {
		// 	ii, _ := strconv.Atoi(string(ss))
		// 	arr = append(arr, ii)
		// }
		input2 := bufio.NewScanner(os.Stdin)
		inputStr2 := input2.Text()
		n, _ = strconv.Atoi(inputStr2)
		// }
		fmt.Println(minKBitFlips(arr, n))
	}

}

func minKBitFlips(A string, K int) int {
	flippedTime, count := 0, 0
	for i := 0; i < len(A); i++ {
		if i >= K && A[i-K] == 2 {
			flippedTime--
		}
		if flippedTime%2 == strconv.A[i] {
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
