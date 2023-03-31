package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestJan1(t *testing.T) {
	tmap := make(map[string]string)
	tmap["1"] = "1"
	fmt.Println(len(tmap))

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Done()

	wgtest := sync.WaitGroup{}
	wgtest.Add(1)
	wgtest.Done()
}
