package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	num := 1
	elem := reflect.ValueOf(&num).Elem()         //variable -> interface variable (value, concrete type) -> reflect.Value
	if canset := elem.CanSet(); canset == true { //addressable
		elem.SetInt(2)
		fmt.Println(num)
	}
	fmt.Println(elem.Interface().(int)) // reflect.Value -> interface variable -> variable   Type Assertion

	a_int, _ := strconv.Atoi("1") //ASCII to integer
	fmt.Println(a_int)

	a_string1 := strconv.Itoa(a_int)
	fmt.Println(a_string1)
	a_string2 := fmt.Sprint(a_int)
	fmt.Println(a_string2)
}
