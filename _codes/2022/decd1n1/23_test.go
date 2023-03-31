package main

import (
	"fmt"
	"testing"
)

func Test72(t *testing.T) {
	fmt.Println(minDistance("ab", "ac"))
}
func minDistance(word1 string, word2 string) int {
	min3 := func(elements ...int) int {
		min := elements[0]
		for _, e := range elements {
			if e < min {
				min = e
			}
		}
		return min
	}

	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min3(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
