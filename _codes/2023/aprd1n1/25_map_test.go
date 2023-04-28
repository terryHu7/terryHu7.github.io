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

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for k, v := range m {
		if k == "two" {
			m["four"] = 4
		}
		fmt.Printf("key=%s, value=%d\n", k, v)
	}
}
