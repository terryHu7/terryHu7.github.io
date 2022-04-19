package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan() //读取一行内容
	m, _ := strconv.Atoi(strings.Split(input.Text(), ",")[0])
	n, _ := strconv.Atoi(strings.Split(input.Text(), ",")[1])
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		input.Scan() //读取一行内容
		for j := 0; j < n; j++ {
			res[i][j], _ = strconv.Atoi(strings.Split(input.Text(), ",")[j])
		}
	}
	fmt.Println(res)
}
