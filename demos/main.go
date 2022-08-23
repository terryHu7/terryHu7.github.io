package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	fmt.Println("Open file...")
	filepath := "/home/hutianyu/0823.txt"

	for count, max := 0, 20; count < max; count++ {
		fd, err := os.OpenFile(filepath, os.O_RDWR, 0)
		if err != nil {
			fmt.Printf("fail to open %s", filepath)
		}
		defer fd.Close()
		content, _ := ioutil.ReadAll(fd)
		fmt.Println(string(content))
		fd.WriteString(fmt.Sprint(count))
		time.Sleep(time.Second)
	}
}
