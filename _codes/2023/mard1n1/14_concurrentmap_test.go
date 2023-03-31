package mard1n1

import (
	"fmt"
	"sync"
	"testing"
)

var testmap map[int]int = map[int]int{}
var testmapLock sync.RWMutex
var wg *sync.WaitGroup = &sync.WaitGroup{}

func TestConcurrentmap(t *testing.T) {
	for i := 0; i < 200; i++ {
		wg.Add(1)
		// go write()
		go read()
	}
	wg.Wait()
}

func write() {
	for i := 0; i < 20; i++ {
		testmapLock.Lock()
		testmap[i] = i
		testmapLock.Unlock()
	}
	wg.Done()
}

func read() {
	for i := 0; i < 20; i++ {
		// testmapLock.RLock()
		fmt.Println(testmap[i])
		// testmapLock.RUnlock()
	}
	wg.Done()
}
