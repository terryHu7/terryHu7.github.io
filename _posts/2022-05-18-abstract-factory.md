---
layout: post
title: "abstract factory <-> go interface"
subtitle: ''
author: "Terry"
header-style: text
tags:
  - Golang
  - Design Pattern
---

> **`A struct` implement `an interface` if all `methods` of the interface have been implemented by the struct.**


```go
package main

import "fmt"

type abstract_factory interface {
	method() (product string)
}

// A struct implement an interface
// if all methods of the interface have been implemented by the struct.
type nike struct {
}

func (nike nike) method() (product string) {
	return `
	_____
  	|  /|          
	|\/ |______   
	|______|_|_|
	`
}

func getFactory(option string) abstract_factory {
	switch option {
	case "nike":
		{
			return nike{}
		}
	default:
		{
			return nike{}
		}
	}

}

func main() {
	nike := getFactory("nike")
	fmt.Println(nike.method())
}
```