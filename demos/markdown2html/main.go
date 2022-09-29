package main

import (
	"bufio"
	"flag"
	"io"
	"os"
	"strings"
)

//Usage: ./markdown2html -inPath "./tmp.txt" -outPath "./newtmp.txt"
var inPath, outPath *string

func init() {
	inPath = flag.String("inPath", "", "")
	outPath = flag.String("outPath", "", "")
}

func main() {
	flag.Parse()
	md2html(*inPath, *outPath)
}

func md2html(inPath, outPath string) {
	infd, err := os.OpenFile(inPath, os.O_RDWR, 0666)
	if err != nil {
		panic("fail to open inPath")
	}
	defer func() {
		infd.Close()
		// os.Remove(inPath)
	}()

	outfd, err := os.OpenFile(outPath, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		panic("fail to open/create outPath")
	}

	r := bufio.NewReader(infd)
	for {
		line, err := r.ReadString('\n')

		if strings.HasPrefix(line, "![](") {
			if strings.HasSuffix(line, "g)\n") || strings.HasSuffix(line, "f)\n") {
				image := line[4 : len(line)-2]
				bytes := []byte("<img src=\"" + image + "\" width=\"400\">" + "\n")
				outfd.WriteString(string(bytes))
			}
		} else {
			outfd.WriteString(line)
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
