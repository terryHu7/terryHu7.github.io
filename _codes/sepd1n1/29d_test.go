package sep_d1n1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestBufioReader(t *testing.T) {
	input := bufio.NewReader(os.Stdin)
	for {
		line, err := input.ReadString('\n')
		fmt.Printf("line: %v", line)
		if err == io.EOF {
			fmt.Println("END")
			break
		}
	}
}
