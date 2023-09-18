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
	parts := strings.Split(input, "@")

	fullSet := parseCharacterSet(parts[0])
	occupiedSet := parseCharacterSet(parts[1])

	result := adjustCharacterCounts(fullSet, occupiedSet)
	fmt.Println(result)
}

func adjustCharacterCounts(fullSet, occupiedSet map[string]int) string {
	result := make(map[string]int)

	for char, count := range fullSet {
		occupiedCount, exists := occupiedSet[char]
		if exists {
			result[char] = count - occupiedCount
		} else {
			result[char] = count
		}
	}

	var outputParts []string
	for char, count := range result {
		outputParts = append(outputParts, fmt.Sprintf("%s:%d", char, count))
	}
	output := strings.Join(outputParts, ",")
	return output
}

func parseCharacterSet(input string) map[string]int {
	charSet := make(map[string]int)
	pairs := strings.Split(input, ",")
	for _, pair := range pairs {
		parts := strings.Split(pair, ":")
		charSet[parts[0]] = parseInt(parts[1])
	}
	return charSet
}

func parseInt(s string) int {
	var n int
	fmt.Sscanf(s, "%d", &n)
	return n
}