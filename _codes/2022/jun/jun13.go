package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "a" {
		talk = "a"
	} else {
		talk = "b"
	}
	return
}

func main() {
	var peo People = &Student{}
	think := "b"
	fmt.Println(peo.Speak(think))
}
