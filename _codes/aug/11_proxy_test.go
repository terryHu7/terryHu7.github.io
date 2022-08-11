package aug

import (
	. "fmt"
	"testing"
)

type Client interface {
	Do() string
}

//
type RealClient struct{}

func (RealClient) Do() string {
	return "real"
}

//
type Proxy struct {
	realclient Client
}

func (p Proxy) Do() string {
	return "pre:" + p.realclient.Do() + ":after"
}

func TestProxy(t *testing.T) {
	proxy := &Proxy{}

	Println(proxy.Do())
}

//really similar to decorator
