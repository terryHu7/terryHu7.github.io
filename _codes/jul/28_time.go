package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(strings.ReplaceAll(now.Format("15:04:05"), ":", "_"))
}
