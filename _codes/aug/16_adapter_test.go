package aug

import (
	"fmt"
	"testing"
)

//plug->covertor->socket

/////////////////
type Adaptee interface {
	AdapteeMethod() string
}

type AdapteeImpl struct{}

func (*AdapteeImpl) AdapteeMethod() string {
	return "adaptee"
}

//////////////////

type Adapter interface {
	AdapterMethod() string
}

type AdapterImpl struct {
	Adaptee
}

func (a *AdapterImpl) AdapterMethod() string {
	return a.AdapteeMethod()
}

func TestAdapter(t *testing.T) {
	adaptee := &AdapteeImpl{}
	adapter := &AdapterImpl{
		Adaptee: adaptee,
	}
	fmt.Println(adapter.AdapterMethod())
}
