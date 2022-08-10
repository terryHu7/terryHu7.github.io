package aug

import (
	"fmt"
	"testing"
)

type Singletonclass struct {
	Property int
}

var OneAndOnlyInstance *Singletonclass

func New() *Singletonclass {
	if OneAndOnlyInstance == nil {
		return &Singletonclass{Property: 728} //set a default, easier to confirm there is only one instance.
	}

	return OneAndOnlyInstance
}

func TestSingleton(t *testing.T) {
	n1 := New()
	n2 := New()
	fmt.Printf("%v,%v\n", &n1, &n2)
}
