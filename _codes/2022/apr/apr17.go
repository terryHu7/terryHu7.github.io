package main

import (
	"fmt"
	"io"
	"strconv"
)

func main() {

	input := 0
	for {
		_, err := fmt.Scanf("%d", &input)
		if err != nil {
			if err == io.EOF {
				break
			}
		} else {
			inputArr := []byte(strconv.Itoa(input))
			result := true
			for i := 0; i < len(inputArr)/2; i++ {
				if inputArr[i] != inputArr[len(inputArr)-1-i] {
					result = false
				}
			}
			fmt.Println(result)
		}
	}
}
