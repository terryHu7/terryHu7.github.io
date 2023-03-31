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
	for input.Scan() {
		temp, _ := strconv.Atoi(strings.Split(input.Text(), " ")[0])
		//下标为字符串按分隔符分割后索引的字符位置
		//每次循环读入一行，并按行将一行数据按分隔进行读入，十分方便
		fmt.Println(temp)

	}
}
