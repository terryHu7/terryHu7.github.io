package mard1n1

import (
	"fmt"
	"os"
	"testing"
)

func TestFilePath(t *testing.T) {
	fullpath := "/home/user/"
	fullPathInfo, err := os.Stat(fullpath)
	if err == nil && fullPathInfo.IsDir() {
		fmt.Println("dir")
		return
	}
	fmt.Println("correct")
}
