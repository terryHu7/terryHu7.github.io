package main

import "fmt"

var frameIdCount = 98

func main() {
	fmt.Println(FrameIdGet(), FrameIdGet(), FrameIdGet())

}

func FrameIdGet() int {
	count := frameIdCount
	frameIdCount = count + 1
	if frameIdCount == 100 {
		frameIdCount++
	}
	return count
}
