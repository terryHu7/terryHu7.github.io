package main

import "fmt"

func main() {
	fmt.Println(2 * (30 / 30))

	testmap1 := make(map[string]*string)
	// testmap2 := make(map[string]string)
	// testmap3 := make(map[string]string)
	fmt.Println(testmap1["null"] == nil)
	var t []interface{}
	t = append(t, 1, 2, 3, "abc")
	fmt.Println(t)
}
