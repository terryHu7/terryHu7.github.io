package aug

import (
	"fmt"
	"testing"
)

type Phone interface {
	Slogan() string
}

//
type Iphone11 struct{}

func (i11 Iphone11) Slogan() string {
	return "Iphone 11"
}

//
type Iphone11Pro struct {
	i11 Phone //**
}

func (i11pro Iphone11Pro) Slogan() string {
	return i11pro.i11.Slogan() + " Pro" //**
}

func TestDecorator(t *testing.T) {
	decorator := Iphone11Pro{Iphone11{}}

	fmt.Println(decorator.Slogan())
}

//really similar to proxy
