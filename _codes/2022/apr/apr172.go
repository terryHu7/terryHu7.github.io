package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		inputstr := input.Text()
		inputArr := []byte(inputstr)
		result := true
		for i := 0; i < len(inputArr)/2; i++ {
			if inputArr[i] != inputArr[len(inputArr)-1-i] {
				result = false
			}
		}
		fmt.Println(result)
	}
}
