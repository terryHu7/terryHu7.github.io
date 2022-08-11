package aug

import (
	"fmt"
	. "fmt"
	"testing"
)

type i1 interface {
	func1() string
}

type s1 struct{}

func (s1) func1() string {
	return "func1"
}

func TestInterface(t *testing.T) {
	struct1 := s1{}
	Printf("%p\n", &struct1)
	Println(struct1.func1())

	var s2 i1 = s1{}
	fmt.Println(s2)

}
