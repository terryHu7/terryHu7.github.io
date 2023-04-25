package main

import (
	"fmt"
	"testing"
)

func TestNQueens(t *testing.T) {
	fmt.Println(totalNQueens(4))
}

var result [][]string

func solveNQueens(n int) [][]string {
	//create 2d array
	blocks := make([][]byte, n)
	for i := 0; i < n; i++ {
		blocks[i] = make([]byte, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			blocks[i][j] = '.'
		}
	}

	curRow := 0
	result = [][]string{}

	var backtrack func(int, int, [][]byte)
	backtrack = func(n int, curRow int, blocks [][]byte) {
		if curRow == n {
			tmp := []string{}
			for _, row := range blocks {
				tmp = append(tmp, string(row))
			}
			result = append(result, tmp)
			return
		}

		for i := 0; i < n; i++ {
			if isValid(curRow, i, blocks) {
				blocks[curRow][i] = 'Q'
				backtrack(n, curRow+1, blocks)
				blocks[curRow][i] = '.'
			}
		}
	}

	backtrack(n, curRow, blocks)
	return result
}

func isValid(curRow int, curCol int, blocks [][]byte) bool {
	for i := 0; i < curRow; i++ {
		if blocks[i][curCol] == 'Q' {
			return false
		}
	}

	for i, j := curRow-1, curCol-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if blocks[i][j] == 'Q' {
			return false
		}
	}
	for i, j := curRow-1, curCol+1; i >= 0 && j < len(blocks); i, j = i-1, j+1 {
		if blocks[i][j] == 'Q' {
			return false
		}
	}

	return true
}

var count int

func totalNQueens(n int) int {
	//create 2d array
	blocks := make([][]byte, n)
	for i := 0; i < n; i++ {
		blocks[i] = make([]byte, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			blocks[i][j] = '.'
		}
	}

	curRow := 0
	count = 0

	var backtrack func(int, int, [][]byte)
	backtrack = func(n int, curRow int, blocks [][]byte) {
		if curRow == n {
			count++
			return
		}

		for i := 0; i < n; i++ {
			if isValid(curRow, i, blocks) {
				blocks[curRow][i] = 'Q'
				backtrack(n, curRow+1, blocks)
				blocks[curRow][i] = '.'
			}
		}
	}

	backtrack(n, curRow, blocks)
	return count
}
