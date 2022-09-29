package main

import (
	"bufio"
	"flag"
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
	md2html(*filePath)
}

func md2html(filePath string) {
	fd, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		panic("wrong filePath")
	}
	defer fd.Close()

	r := bufio.NewReader(fd)
	pos := int64(0)
	for {
		line, err := r.ReadString('\n')

		if strings.HasPrefix(line, "![](") && strings.HasSuffix(line, "g)\n") {
			image := line[4 : len(line)-2]
			bytes := []byte("<img src=\"" + image + "\" width=\"400\">" + "\n")
			fd.WriteAt(bytes, pos)
			pos += (int64(len(bytes)))
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
