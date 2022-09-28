package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var filePath *string

func init() {
	filePath = flag.String("filePath", "", "")
}

func main() {
	flag.Parse()

	fd, err := os.OpenFile(*filePath, os.O_RDWR, 0666)
	if err != nil {
		panic("wrong filePath")
	}
	defer fd.Close()

	r := bufio.NewReader(fd)
	pos := int64(0)
	for {
		line, err := r.ReadString('\n')
		fmt.Printf("line: %v", line)
		if strings.HasPrefix(line, "![](") && strings.HasSuffix(line, "g)\n") {
			image := line[4 : len(line)-2]
			fmt.Printf("image: %v\n", image)
			bytes := []byte("<img src=\"" + image + "\" width=\"400\">" + "\n")
			fmt.Printf("pos: %v\n", pos)
			fd.WriteAt(bytes, pos)
			pos += (int64(len(bytes) - len(line)))
		} else {
			pos += int64(len(line))
		}

		// cmd := exec.Command("sed", "-r", "-i", "")
		// error := cmd.Run()
		// if error != nil {
		// }

		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
}
