package main

import "fmt"

var frameIdCount = 98
var bgdId int = 65535

func main() {
	// fmt.Println(FrameIdGet(), FrameIdGet(), FrameIdGet())
	fmt.Println(fmt.Sprint(bgdId))

}

func FrameIdGet() int {
	count := frameIdCount
	frameIdCount = count + 1
	if frameIdCount == 100 {
		frameIdCount++
	}
	return count
}
