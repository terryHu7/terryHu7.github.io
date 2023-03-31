package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var uniqueid int

func TestTime(t *testing.T) {
	fmt.Println(getUniqueString())
	fmt.Println(getUniqueString())
	fmt.Println(getUniqueString())

	fmt.Println(strconv.Atoi(""))
	a := false
	fmt.Println(func() bool {
		if a {
			return true
		} else {
			return false
		}
	}())
}

func getUniqueString() string {
	uniquestr := fmt.Sprintf("%s_%s", time.Now().Format("0201_1504"), fmt.Sprint(uniqueid))
	uniqueid++
	return uniquestr
}
