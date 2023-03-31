package aug

import (
	"fmt"
	"testing"
)

//plug->covertor->socket

/////////////////
type Socket interface {
	Charge() string
}

type TwoHoleSocket struct{}

func (*TwoHoleSocket) Charge() string {
	return "Charging..."
}

//////////////////

type Convertor interface {
	Convert() string
}

type MyConvertor struct {
	Socket
}

func (m *MyConvertor) Convert() string {
	return m.Charge()
}

func TestAdapter(t *testing.T) {
	convertor := &MyConvertor{
		Socket: &TwoHoleSocket{}, //covertor->socket
	}
	fmt.Println(convertor.Convert()) //plug->covertor
}
