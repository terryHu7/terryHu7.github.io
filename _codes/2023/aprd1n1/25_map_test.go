package main

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	tmap := map[int]int{}
	tmap[1] = 1
	tmap[2] = 2
	for _, t := range tmap {
		fmt.Println(t)
		if t == 2 {
			tmap[3] = 3
		}
	}
}
