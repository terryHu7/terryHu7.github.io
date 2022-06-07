package main

import "fmt"

func main() {
	fmt.Println(0&1 == 0)

	map1 := map[string]string{}
	handler1(map1)
	fmt.Println(map1)
	handler2(map1)
	fmt.Println(map1)

}

func handler1(map2 map[string]string) {
	map2["test"] = "test" //no need to return
}

func handler2(map2 map[string]string) {
	map2 = map[string]string{"hhh": "hhh"} //another map
}
