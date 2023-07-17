package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestTT(t *testing.T) {
	rand.Seed(2)

	randomNumber := rand.Intn(5)
	randomNumber2 := rand.Intn(5)
	randomNumber3 := rand.Intn(5)
	fmt.Println(randomNumber, randomNumber2, randomNumber3)
}
