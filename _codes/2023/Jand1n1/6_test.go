package main

import (
	"fmt"
	"testing"
)

type human struct {
	Name string
	Age  int
}

func newinterface() interface{} {
	return &human{Name: "Terry", Age: 24}
}

func TestJan6(t *testing.T) {
	sample1 := newinterface()
	fmt.Printf("%+v\n", sample1)
	sample2, ok := newinterface().(*human)
	fmt.Printf("%+v,%+v\n", sample2.Name, ok)

}
