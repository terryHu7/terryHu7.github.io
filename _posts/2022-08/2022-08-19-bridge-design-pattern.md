---
layout: post
title: "Bridge"
subtitle: ''
author: "Terry"
header-style: text
tags:
  - Design Pattern
---

![bridge](http://www.17bigdata.com/study/static.codebaoku.com/pics/85/e9/85e923603fa9520cbcaa02efbb4bb214.jpg)

```go
package aug

import (
	. "fmt"
	"reflect"
	"testing"
)

type Computer interface {
	SetNprint()
}

type Mac struct {
	printer Printer //Bridge: interface, not like HP/Epison
}

func (m *Mac) SetNprint(p Printer) {
	m.printer = p
	Printf("Build a bridge between %v and %v...\n", reflect.TypeOf(m).Elem().Name(), reflect.TypeOf(p).Elem().Name())
	p.Print()
}

type Win struct {
	printer Printer
}

func (m *Win) SetNprint(p Printer) {
	m.printer = p
	p.Print()
}

//////////////////////////////////////////////////
type Printer interface {
	Print()
}

type HP struct{}

func (hp *HP) Print() {
	Println("Printing by HP...")
}

type Epison struct{}

func (epison *Epison) Print() {
	Println("Printing by Epison...")
}

func TestBridge(t *testing.T) {
	mac := &Mac{}
	hp := &HP{}
	epison := &Epison{}
	mac.SetNprint(hp)
	mac.SetNprint(epison)
}

//http://www.17bigdata.com/study/programming/it-go/it-go-234126.html


```