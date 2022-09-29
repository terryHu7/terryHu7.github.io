package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	for {
		line, err := input.ReadString('\n')
		s := strings.Split(string(line[0:len(line)-1]), ",") // '\n' does not count
		nums := []int{}
		for i := 0; i < len(s); i++ {
			num, _ := strconv.Atoi(s[i])
			fmt.Printf("num: %v\n", num)
			//append
		}
		if err == io.EOF {
			fmt.Println("END")
			break
		}
	}
}
