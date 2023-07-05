package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestHeap(t *testing.T) {
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(len([]int{0, 1, 2, 3, 4, 5})))
}
