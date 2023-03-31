package aug

import (
	"fmt"
	"testing"
)

func HeapSort(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		sink(a, i, len(a))
	}
	for i := len(a) - 1; i >= 1; i-- {
		a[0], a[i] = a[i], a[0]
		sink(a, 0, i)
	}
	return a
}
func sink(a []int, i int, length int) {
	for {
		l := i*2 + 1
		r := i*2 + 2
		idx := i
		if l < length && a[l] > a[idx] {
			idx = l
		}
		if r < length && a[r] > a[idx] {
			idx = r
		}
		if idx == i {
			break
		}
		a[i], a[idx] = a[idx], a[i]
		i = idx
	}
}

func TestHeapSort(t *testing.T) {
	fmt.Println(HeapSort([]int{3, 7, 2, 4, 9, 10, 6}))
}
