package aug

import (
	"fmt"
	"testing"
)

//simply speaking, it's "for"
type Iterator26 interface {
	NewIterator(m iterSet) Iter
}

type Iter struct {
	C chan interface{}
}

func newIter(i *iterSet) Iter {
	iter := Iter{make(chan interface{})}

	go func() {
		for k := range i.m {
			iter.C <- k
		}
		close(iter.C)
	}()

	return iter
}

type iterSet struct {
	m map[string]bool
}

func (i *iterSet) Add(k string) {
	i.m[k] = true
}

func (i *iterSet) NewIterator() Iter {
	return newIter(i)
}

func TestIterator(t *testing.T) {
	aSet := iterSet{map[string]bool{}}
	aSet.Add("hello")
	aSet.Add("world")

	iter := aSet.NewIterator()

	for v := range iter.C {
		fmt.Printf("key: %s\n", v.(string))
	}
}
