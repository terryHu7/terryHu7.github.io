package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Split(input, ",")

	nestedMap := make(map[int][]string)
	maxLevel := parseInt(parts[0])

	for i := 1; i < len(parts)-1; i += 2 {
		word := parts[i]
		level := parseInt(parts[i+1])
		nestedMap[level] = append(nestedMap[level], word)
	}

	for level := 1; level <= maxLevel; level++ {
		words := nestedMap[level]
		fmt.Println(strings.Join(words, " "))
	}
}

func parseInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}