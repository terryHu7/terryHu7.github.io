// package main

// import (
// 	"fmt"
// )

// func main() {
// 	//input
// 	m, n := 0, 0
// 	r1, c1 := 0, 0
// 	r2, c2 := 0, 0

// 	fmt.Scanf("%d", &m)
// 	fmt.Scanf("%d", &n)
// 	fmt.Scanf("%d", &r1)
// 	fmt.Scanf("%d", &c1)
// 	fmt.Scanf("%d", &r2)
// 	fmt.Scanf("%d", &c2)

// 	output := 0

// 	grid := make([][]int, m)
// 	for i := range grid {
// 		grid[i] = make([]int, n)
// 	}
// 	if r1 >= 0 && r1 < m && c1 >= 0 && c1 < n {
// 		grid[r1][c1] = 1
// 	}
// 	if r2 >= 0 && r2 < m && c2 >= 0 && c2 < n {
// 		grid[r2][c2] = 1
// 	}

// 	visited := make([][]bool, m)
// 	for i := range grid {
// 		visited[i] = make([]bool, n)
// 	}

// 	//
// 	for {
// 		andOne := false
// 		for i := 0; i < m; i++ {
// 			for j := 0; j < n; j++ {
// 				if grid[i][j] == 1 {
// 					diffusion(grid, i, j, &andOne, visited)
// 				}
// 			}
// 		}
// 		if andOne {
// 			output++
// 			for k := 0; k < m; k++ {
// 				for l := 0; l < n; l++ {
// 					visited[k][l] = false
// 				}
// 			}
// 		} else {
// 			break
// 		}
// 	}

// 	//output
// 	fmt.Println(output)
// }

// func diffusion(grid [][]int, i, j int, andOne *bool, visited [][]bool) {
// 	m, n := len(grid), len(grid[0])
// 	if i < 0 || j < 0 || i >= m || j >= n {
// 		return
// 	}

// 	if visited[i][j] == true {
// 		return
// 	}

// 	visited[i][j] = true
// 	if i-1 >= 0 && i-1 < m && grid[i-1][j] == 0 {
// 		grid[i-1][j] = 1
// 		visited[i-1][j] = true
// 		*andOne = true
// 	}

// 	if i+1 >= 0 && i+1 < m && grid[i+1][j] == 0 {
// 		grid[i+1][j] = 1
// 		visited[i+1][j] = true
// 		*andOne = true
// 	}
// 	if j-1 >= 0 && j-1 < n && grid[i][j-1] == 0 {
// 		grid[i][j-1] = 1
// 		visited[i][j-1] = true
// 		*andOne = true
// 	}
// 	if j+1 >= 0 && j+1 < n && grid[i][j+1] == 0 {
// 		grid[i][j+1] = 1
// 		visited[i][j+1] = true
// 		*andOne = true
// 	}
// 	return
// }
