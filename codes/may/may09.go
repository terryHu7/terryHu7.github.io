package main

import "fmt"

type human struct {
	id  int
	sex bool
}

type man struct {
	human
	sex int
}

func (h *human) gender() {
	fmt.Println(h.sex)
}
func (m *man) gender() { //override
	fmt.Println(m.sex)
}

func main() {
	man1 := man{
		human{
			id:  320323,
			sex: true},
		1}

	man1.human.gender() //anti-override
	man1.gender()
}
