package main

import (
	"fmt"
	"testing"
)

func TestSudoku(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)

	// tmap := map[int]int{}
	// tmap[1] = 1
	// tmap[2] = 2
	// for _, t := range tmap {
	// 	fmt.Println(t)
	// 	if t == 2 {
	// 		tmap[3] = 3
	// 	}
	// }

}

func solveSudoku(board [][]byte) {
	vertices := [][]int{
		{0, 0},
		{0, 3},
		{0, 6},
		{3, 0},
		{3, 3},
		{3, 6},
		{6, 0},
		{6, 3},
		{6, 6},
	}
	leftNumbersOfBlocks := make([]map[byte]byte, 9)
	for i := 0; i < len(leftNumbersOfBlocks); i++ {
		leftNumbersOfBlocks[i] = map[byte]byte{}
		for j := 0; j < 9; j++ {
			leftNumbersOfBlocks[i][byte(j+1)+'0'] = byte(j+1) + '0'
		}
	}

	for index := 0; index < 9; index++ {
		startRow, startCol := vertices[index][0], vertices[index][1]
		for i := startRow; i < startRow+3; i++ {
			for j := startCol; j < startCol+3; j++ {
				if board[i][j] == '.' {
					continue
				}
				b := board[i][j]
				delete(leftNumbersOfBlocks[index], b)
			}
		}
	}

	backtracking(vertices, leftNumbersOfBlocks, board)
	fmt.Println(board)
}

func isvalid(row int, col int, k byte, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == k {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if board[i][col] == k {
			return false
		}
	}

	// startrow := (row / 3) * 3
	// startcol := (col / 3) * 3
	// for i := startrow; i < startrow+3; i++ {
	// 	for j := startcol; j < startcol+3; j++ {
	// 		if board[i][j] == k {
	// 			return false
	// 		}
	// 	}
	// }
	return true
}

func backtracking(vertices [][]int, leftNumbersOfBlocks []map[byte]byte, board [][]byte) bool {
	for index := 0; index < 9; index++ {
		startRow, startCol := vertices[index][0], vertices[index][1]
		for i := startRow; i < startRow+3; i++ {
			for j := startCol; j < startCol+3; j++ {
				if board[i][j] != '.' {
					continue
				}

				tmp := deepCopyMap(leftNumbersOfBlocks[index])
				for _, b := range tmp {
					if isvalid(i, j, b, board) {
						board[i][j] = b
						delete(leftNumbersOfBlocks[index], b)

						if backtracking(vertices, leftNumbersOfBlocks, board) {
							return true
						}

						board[i][j] = '.'
						leftNumbersOfBlocks[index][b] = b
					}
				}
				return false
			}
		}
	}
	return true
}

func deepCopyMap(original map[byte]byte) map[byte]byte {
	copied := make(map[byte]byte, len(original))

	for key, value := range original {
		copied[key] = value

	}
	return copied
}
