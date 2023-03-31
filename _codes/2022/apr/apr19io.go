package main

import (
	"fmt"
	"unicode"
)

func main() {
	// r := bufio.NewReader(os.Stdin)
	// input, _ := r.ReadString('\n')
	// fmt.Println(input[0])
	r := []rune(`Hi1`)
	fmt.Println(unicode.IsLetter(r[2]))
	fmt.Println(unicode.IsDigit(r[2]))
	fmt.Println(unicode.IsLower(r[1]))
}
