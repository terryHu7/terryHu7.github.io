package main

import (
	"fmt"
)

func main() {
	m := map[string]string{}
	fmt.Println(m == nil)
	fmt.Println(len(m) == 0)
	fmt.Println(m["123"])
	fmt.Println(m["123"] == "")
	if v, ok := m["123"]; ok {
		fmt.Println("ok", v)
	} else {
		fmt.Println("not ok", v)
	}
	// req := []byte(`{
	// 	"code":	200,
	// 	"message":	"Success"
	// }`)
	// reqmap := map[string]interface{}{}
	// _ = json.Unmarshal(req, &reqmap)
	// fmt.Println(fmt.Sprint(reqmap["code"]) == "200")
	// fmt.Printf("%v,%T", reqmap["code"], reqmap["code"])
}
