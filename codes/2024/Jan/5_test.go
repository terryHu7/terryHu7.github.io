package Jan

import (
	"fmt"
	"testing"
)

func TestUnicode(t *testing.T) {
	str := "国"
	utf8Bytes := []byte(str)
	for _, b := range utf8Bytes {
		fmt.Printf("%x ", b) // 以二进制形式打印每个字节
	}

	var b byte = 78
	fmt.Println(string(b))
}
